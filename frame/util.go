package frame

import (
	"context"
	"crypto/tls"
	"fmt"
	//log "iceberg/frame/icelog"
	log "iceberg/frame/mantlog"
	"math/rand"
	"net"
	"net/mail"
	"net/smtp"
	"net/textproto"
	"strings"
	"syscall"
	"time"
)

// 定义格式化时间格式
const (
	Normalformat = "2006-01-02 15:04:05"
)

// ValueFrom 从Context获取KEY Value对
func ValueFrom(ctx context.Context, key string) string {
	k := textproto.CanonicalMIMEHeaderKey(key)
	v, ok := ctx.Value(k).(string)
	if ok {
		return v
	}
	return ""
}

// Netip 返回内网IP地址
func Netip() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err.Error())
	}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() &&
			ipnet.IP.To4() != nil {
			return strings.Split(addr.String(), "/")[0]
		}
	}
	return ""
}

// RandPort 随机端口 1024 - 65535
func RandPort() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		port := r.Intn(64511) + 1024
		sa := new(syscall.SockaddrInet4)
		sa.Port = port
		sa.Addr = [4]byte{0x7f, 0x00, 0x00, 0x01}
		if s, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, syscall.IPPROTO_TCP); err != nil {
			log.Debug("Socket:", err.Error())
			continue
		} else if err := syscall.Bind(s, sa); err != nil {
			log.Debug("端口被占用:", port)
			continue
		} else {
			syscall.Close(s)
			log.Debugf("端口%d可用", sa.Port)
			return fmt.Sprintf("%d", port)
		}
	}
}

// ShapingTime 把超时时间整形,得到的超时时间显现有规律的间隔; 参数:
// cell time.Duration 超时的时间粒度
// ticks int64 超时的时间滴答数。
func ShapingTime(begin time.Time, cell time.Duration, ticks int64) time.Time {
	a := begin.UnixNano() / int64(cell)
	b := (a + 1 + ticks) * int64(cell)
	return time.Unix(0, b)
}

// Alarm Alarm
func Alarm(srvName, toEmail string) {
	warnTime := time.Now().Format("2006-01-02 15:04:05")
	var alarmTitle = fmt.Sprintf("【注意】, 进程[%s]于 %s 启动", srvName, warnTime)
	var alarmContent = fmt.Sprintf("【注意】, 进程[%s@%s]于 %s 启动。请留意该现象是否正常。",
		srvName, Instance().localListenAddr, warnTime)
	AlarmEvent(srvName, toEmail, alarmTitle, alarmContent)
}

// AlarmEvent 程序启动告警
func AlarmEvent(srvName, toEmail, title, content string) {
	if srvName == "" || toEmail == "" {
		return
	}

	var from = mail.Address{
		Name:    "冰人",
		Address: "no-reply@laoyuegou.com"}

	var to = mail.Address{
		Name:    "",
		Address: toEmail}

	var msg = []byte("From: \"alarmrobot\"<" + from.Address + ">\r\n" +
		"To: " + toEmail + "\r\n" +
		"Subject: " + title + "\r\n" +
		"\r\n" +
		content + "\r\n")

	var serverHost = "smtp.exmail.qq.com:465"
	host, _, _ := net.SplitHostPort(serverHost)
	auth := smtp.PlainAuth("", from.Address, "AsdF1234", host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", serverHost, tlsconfig)
	if err != nil {
		log.Error(err)
		return
	}
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Error(err)
		return
	}

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			// Auth
			if err = c.Auth(auth); err != nil {
				log.Error(err)
				return
			}
		}
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Error(err)
		return
	}
	if err = c.Rcpt(to.Address); err != nil {
		log.Error(err.Error())
		return
	}
	// Data
	w, err := c.Data()
	if err != nil {
		log.Error(err)
		return
	}
	_, err = w.Write(msg)
	if err != nil {
		log.Error(err)
		return
	}
	err = w.Close()
	if err != nil {
		log.Error(err)
		return
	}
	c.Quit()
}
