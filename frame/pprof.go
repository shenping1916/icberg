package frame

import (
	"fmt"
	//log "iceberg/frame/icelog"
	log "iceberg/frame/mantlog"
	"os"
	"runtime/pprof"
	"time"
)

// profile responds with the pprof-formatted cpu profile.
func profile(cmd string) {

	if cmd != "cpu" && cmd != "memory" {
		log.Warnf("bad cmd %s", cmd)
		return
	}
	log.Debug("开始Profile:", cmd)
	f, err := os.Create(fmt.Sprintf("%s.pprof", cmd))
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer f.Close()

	switch cmd {
	case "cpu":
		if err = pprof.StartCPUProfile(f); err != nil {
			log.Error(err.Error())
			return
		}
		defer pprof.StopCPUProfile()
	case "memory":
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Error(err.Error())
			return
		}
	}
	// 等待60s收集信息
	time.Sleep(time.Second * 60)
	log.Debug("结束Profile:", cmd)
}
