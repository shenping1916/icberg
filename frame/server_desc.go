package frame

import (
	"iceberg/frame/protocol"

	objectid "github.com/nobugtodebug/go-objectid"
)

type methodHandler func(srv interface{}, ctx Context) error

// Authorization 认证类型
type Authorization string

// 定义 Authorization 值
const (
	Stream   Authorization = "stream"   // 服务间流式调用，只能服务间相互调用的接口，内部服务可以调用任意体系内接口，而steam接口只能服务间调用
	Internal Authorization = "internal" // 只允许内网访问的接口
	App      Authorization = "app"      // 只允许App访问的接口
	Openapi  Authorization = "openapi"  // 只允许第三方访问的接口

)

// MethodDesc 代表RPC方法的描述
type MethodDesc struct {
	// 认证
	A Authorization

	// 方法名称
	MethodName string

	// 调起方法的句柄
	Handler methodHandler
}

// ServiceDesc 服务描述
type ServiceDesc struct {
	Version     string
	ServiceName string
	// 指向具体服务
	// 并会检查服务是否实现了proto文件描述的方法
	HandlerType interface{}
	Methods     []MethodDesc
	Metadata    interface{}
	ServiceURI  []string
}

// ReadyTask 准备请求的任务
func ReadyTask(fc Context,
	srvMethod, srvName, srvVersion string,
	in interface{}, opts ...CallOption) (*protocol.Proto, error) {

	c := defaultCallInfo()
	for _, o := range opts {
		if err := o.before(c); err != nil {
			return nil, err
		}
	}

	var task protocol.Proto
	if fc.Bizid() == "" {
		task.Bizid = objectid.New().String()
	} else {
		task.Bizid = fc.Bizid()
	}

	task.ServeMethod = srvMethod
	task.Format = c.format

	task.Header = make(map[string]string)
	for k := range c.header {
		task.Header[k] = c.header.Get(k)
	}

	task.Form = make(map[string]string)
	for k, v := range c.form {
		task.Form[k] = v
	}

	// 默认序列化方式为JSON
	if task.Format == protocol.RestfulFormat_FORMATNULL {
		task.Format = protocol.RestfulFormat_JSON
	}

	task.RequestID = GetInnerID()
	task.ServeURI = "/services/" + srvVersion + "/" + srvName
	task.Method = protocol.RestfulMethod_POST

	b, err := protocol.Pack(task.Format, in)
	if err != nil {
		return nil, err
	}

	task.Body = b
	return &task, nil
}
