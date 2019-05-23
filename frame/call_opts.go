package frame

import (
	"iceberg/frame/protocol"
	"net/http"
)

type callInfo struct {
	maxReceiveMessageSize *int // TODO
	maxSendMessageSize    *int // TODO
	failFast              bool
	ignore                bool
	form                  map[string]string
	format                protocol.RestfulFormat
	header                http.Header
}

// CallOption 请求Option
type CallOption interface {
	before(*callInfo) error
	after(*callInfo)
}

func defaultCallInfo() *callInfo {
	return &callInfo{failFast: true}
}

type beforeCall func(c *callInfo) error

func (o beforeCall) before(c *callInfo) error { return o(c) }
func (o beforeCall) after(c *callInfo)        {}

type afterCall func(c *callInfo) error

func (o afterCall) before(c *callInfo) error { return nil }
func (o afterCall) after(c *callInfo)        { o(c) }

// Header 附加Header信息去请求
func Header(header http.Header) CallOption {
	return beforeCall(func(c *callInfo) error {
		c.header = header
		return nil
	})
}

// Format 请求数据序列化方式
func Format(format protocol.RestfulFormat) CallOption {
	return beforeCall(func(c *callInfo) error {
		c.format = format
		return nil
	})
}

// Form With 如果请求参数是Query或者Form则需在Option中使用此接口函数
func Form(f map[string]string) CallOption {
	return beforeCall(func(c *callInfo) error {
		c.form = f
		return nil
	})
}

// Ignore 忽略响应的自动解析
func Ignore(is bool) CallOption {
	return afterCall(func(c *callInfo) error {
		c.ignore = is
		return nil
	})
}
