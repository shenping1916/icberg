package main

import (
	"iceberg/frame"
	"os"

	"iceberg/frame/config"
	//log "iceberg/frame/icelog"
	log "iceberg/frame/mantlog"
	"time"
)

// Config 配置
type Config struct {
	IP   string         `json:"ip"`
	Etcd config.EtcdCfg `json:"etcdCfg"`
	Port string         `json:"prot"`
}

// S1 对象
type S1 struct {
}

// Third 第三方访问测试
func (s *S1) Third(c frame.Context) error {
	return c.String("success")
}

// Internal 内网访问测试
func (s *S1) Internal(c frame.Context) error {
	return c.String("success")
}

// App app 权限访问
func (s *S1) App(c frame.Context) error {
	return c.String("success")
}

// Timeout 超时测试
func (s *S1) Timeout(c frame.Context) error {
	c.Info("receiver time out request....")
	time.Sleep(time.Second * 30)
	return c.String("success")
}

// Stop Stop
func (s *S1) Stop(sig os.Signal) bool {
	log.Infof("hi graceful exit.")
	return true
}
