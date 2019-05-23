package mantlog

import (
	"github.com/shenping1916/mant/log"
)

var Logger *log.Logger

type logger interface {
	GetEnv() string
	GetPath() string
	GetRotate() bool
	GetDaily() bool
	GetCompress() bool
	GetMaxLines() int64
	GetMaxSize() int64
	GetMaxKeepDays() int
}
