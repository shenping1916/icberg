package balance

import (
	"context"
	//log "iceberg/frame/icelog"
	log "iceberg/frame/mantlog"
	"net"
	"sync"
	"sync/atomic"
)

const (
	BALANCE_ROUNDROBIN = "roundrobin"
)

var (
	mb = make(map[string]Balancer)
)

// Balancer 负载均衡器
type Balancer interface {
	Pick([]*RuntimeState) (*net.TCPAddr, error)
	Name() string
}

// BalanceMgr 节点管理
type BalanceMgr struct {
	nLocker                sync.RWMutex
	nlsq                   nodeListSeq
	ring                   map[uint32]*RuntimeState
	runtimeStateReqSuccess chan uint32
	runtimeStateReqFail    chan uint32
	runtimeStatePool       *sync.Pool
}

// Manager new BalanceMgr
func Manager(ctx context.Context) *BalanceMgr {
	mgr := &BalanceMgr{
		ring: make(map[uint32]*RuntimeState),
		runtimeStateReqSuccess: make(chan uint32, 128),
		runtimeStateReqFail:    make(chan uint32, 128),
	}

	mgr.runtimeStatePool = &sync.Pool{
		New: func() interface{} {
			return []*RuntimeState{}
		}}
	go mgr.watcher(ctx)
	return mgr
}

// Register register balance
func Register(bls Balancer) {
	mb[bls.Name()] = bls
}

// RoundRobin 返回roundrobin负载策略对象
func RoundRobin() Balancer {
	return Get(BALANCE_ROUNDROBIN)
}

// Get get balance
func Get(name string) Balancer {
	if bls, ok := mb[name]; ok {
		return bls
	}
	return nil
}

// Add 从ETCD检测到新的实例节点注册，把加入到集群中
func (mgr *BalanceMgr) Add(svrAddr string) {
	hashed := hash([]byte(svrAddr))

	if _, found := mgr.ring[hashed]; found {
		log.Warnf("Hash crash, mgr node [%s:%d] is existed in ring.", svrAddr, hashed)
		return
	}

	tcpAddr, err := net.ResolveTCPAddr("tcp", svrAddr)
	if err != nil {
		log.Errorf("Add a new node %s into hash ring fail,detail=%s", svrAddr, err.Error())
		return
	}
	var rts = &RuntimeState{
		ID:      hashed,
		SvrAddr: tcpAddr,
	}

	mgr.nLocker.Lock()
	mgr.ring[hashed] = rts
	mgr.nlsq.Insert(hashed)
	mgr.nLocker.Unlock()
	log.Debugf("Add a new node %s into hash ring,hashed=%d", svrAddr, hashed)
}

// Remove 检测到新的实例节点删除，从集群中摘除
func (mgr *BalanceMgr) Remove(svrAddr string) string {
	hashed := hash([]byte(svrAddr))
	if v, found := mgr.ring[hashed]; !found {
		log.Warnf("Can't remove node [%s], because the node is not exist.", svrAddr)
		return ""
	} else {
		mgr.nLocker.Lock()
		delete(mgr.ring, hashed)
		if !mgr.nlsq.Remove(hashed) {
			log.Warn("The node [%s] is not exist in nodelist, but exist in ring, Data is not consistent!!!", svrAddr)
		}
		mgr.nLocker.Unlock()
		return v.SvrAddr.String()
	}
}

// Len 节点长度
func (mgr *BalanceMgr) Len() int {
	return len(mgr.nlsq)
}

// Clear 清除所有节点
func (mgr *BalanceMgr) Clear() {
	mgr.nLocker.Lock()
	mgr.nlsq = nodeListSeq{}
	mgr.ring = make(map[uint32]*RuntimeState)
	mgr.nLocker.Unlock()
}

// AllNodeAddr 获取所有节点地址
func (mgr *BalanceMgr) AllNodeAddr() []string {
	mgr.nLocker.Lock()
	var s = make([]string, mgr.Len())
	for i, v := range mgr.nlsq {
		s[i] = mgr.ring[v].SvrAddr.String()
	}
	mgr.nLocker.Unlock()
	return s
}

// ReqSuccess 请求成功
func (mgr *BalanceMgr) ReqSuccess(id uint32) {
	mgr.runtimeStateReqSuccess <- id
}

// ReqFail 请求成功
func (mgr *BalanceMgr) ReqFail(id uint32) {
	mgr.runtimeStateReqFail <- id
}

// Pick 根据负载策略选择一个节点
func (mgr *BalanceMgr) Pick(blr Balancer) (*net.TCPAddr, error) {
	var ringList = mgr.runtimeStatePool.Get().([]*RuntimeState)
	ringList = make([]*RuntimeState, mgr.Len())
	mgr.nLocker.RLock()
	for i, v := range mgr.nlsq {
		ringList[i] = mgr.ring[v]
	}
	mgr.nLocker.RUnlock()
	addr, err := blr.Pick(ringList)
	mgr.runtimeStatePool.Put(ringList)
	return addr, err
}

func (mgr *BalanceMgr) watcher(ctx context.Context) {
	for {
		select {
		case id := <-mgr.runtimeStateReqSuccess:
			mgr.nLocker.RLock()
			if s, ok := mgr.ring[id]; ok {
				atomic.AddUint64(&s.SuccNo, 1)
				atomic.AddUint64(&s.ReqNo, 1)
			}
			mgr.nLocker.RUnlock()

		case id := <-mgr.runtimeStateReqFail:
			mgr.nLocker.RLock()
			if s, ok := mgr.ring[id]; ok {
				atomic.AddUint64(&s.FailNo, 1)
				atomic.AddUint64(&s.ReqNo, 1)
			}
			mgr.nLocker.RUnlock()

		case <-ctx.Done():
			return
		}
	}
}
