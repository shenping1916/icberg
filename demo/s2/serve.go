package main

import (
	"iceberg/frame"
	//log "iceberg/frame/icelog"
	log "iceberg/frame/mantlog"
	"os"
)

// Hi 对象
type Hi struct {
}

// SayHi handel message 01
func (id *Hi) SayHi(c frame.Context) error {
	return c.String("success")
}

// Stop Stop
func (id *Hi) Stop(s os.Signal) bool {
	log.Infof("hi graceful exit.")
	return true
}
