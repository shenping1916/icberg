package balance

import (
	"net"
	"testing"
)

func TestSnapshot(t *testing.T) {
	var nls = nodeListSeq{1, 2, 3, 4, 5}
	snapshot := nls.Snapshot()
	if &snapshot[0] == &nls[0] || &snapshot == &nls {
		t.Error("snapshot fail")
	} else {
		t.Log("snapshot ok:", snapshot)
	}
}

func TestTCPAddr(t *testing.T) {
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:9090")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(addr.Network(), "  ", addr.String())
}

func TestRoundRobin(t *testing.T) {

}
