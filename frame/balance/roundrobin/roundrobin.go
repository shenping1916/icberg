package roundrobin

import (
	"errors"
	"iceberg/frame/balance"
	"net"
	"sync"
	"sync/atomic"
)

var (
	rdrInstance *Roundrobin
	rdrOnce     sync.Once
)

var ErrInstanceNotFound = errors.New("instances not found")

// Roundrobin roundrobin 负载策略
type Roundrobin struct {
	count int32
}

// Pick pick instance
func (rdr *Roundrobin) Pick(
	rtss []*balance.RuntimeState) (*net.TCPAddr, error) {
	if len(rtss) == 0 {
		return nil, ErrInstanceNotFound
	}
	var next = atomic.AddInt32(&rdr.count, 1)
	sc := rtss[next%int32(len(rtss))]
	return sc.SvrAddr, nil
}

// newBalancer build Roundrobin balance
func newBalancer() balance.Balancer {
	rdrOnce.Do(func() {
		rdrInstance = &Roundrobin{}
	})
	return rdrInstance
}

// Build build Roundrobin balance
func (rdr *Roundrobin) Name() string {
	return balance.BALANCE_ROUNDROBIN
}

func init() {
	balance.Register(newBalancer())
}
