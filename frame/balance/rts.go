package balance

import (
	"net"
)

// RuntimeState 运行时状态
type RuntimeState struct {
	ID      uint32
	SvrAddr *net.TCPAddr
	ReqNo   uint64
	FailNo  uint64
	SuccNo  uint64
	Weight  int64
}
