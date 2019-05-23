package balance

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"sort"

	//log "iceberg/frame/icelog"
	log "iceberg/frame/mantlog"
)

type nodeListSeq []uint32

func (h *nodeListSeq) Len() int           { return len(*h) }
func (h *nodeListSeq) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *nodeListSeq) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *nodeListSeq) Insert(x interface{}) {
	// NOTE 可优化点，不直接append, 找到新节点的位置对后面的节点做移位后再插入
	// 这样就不用每次插入都做一次排序了
	*h = append(*h, x.(uint32))
	sort.Sort(h)
	log.Debug("insert node successful.")
}

func (h *nodeListSeq) Remove(x interface{}) bool {
	if h.Len() == 0 {
		return false
	}

	i := sort.Search(h.Len(), func(i int) bool { return (*h)[i] >= x.(uint32) })
	if i < h.Len() && (*h)[i] == x.(uint32) {
		log.Debugf("remove node from nodeList:%d len:%d", i, h.Len())
		*h = append((*h)[:i], (*h)[i+1:]...)
		return true
	} else {
		return false
	}
}

// Snapshot 当前节点快照
func (h *nodeListSeq) Snapshot() nodeListSeq {
	var snapshot = make([]uint32, len(*h))
	copy(snapshot, *h)
	return snapshot
}

func hash(key []byte) uint32 {
	md5Inst := md5.New()
	md5Inst.Write(key)
	result := md5Inst.Sum([]byte(""))

	var value uint32
	buf := bytes.NewBuffer(result)
	err := binary.Read(buf, binary.LittleEndian, &value)
	if err != nil {
		log.Errorf("Calculate [%s] hash failed!", string(key))
	}
	return value
}
