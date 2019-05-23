package mantlog

import (
	"github.com/shenping1916/mant/log"
)

// 日志注册
func RegisterLogger(l logger) {
	Logger = log.NewLogger(4, log.LEVELDEBUG)
	Logger.SetFlag()
	Logger.SetAsynChronous()

	env := l.GetEnv()
	if env != "stage" || env != "prod" {
		Logger.SetOutput(log.CONSOLE)
		Logger.SetColour()
	}
	Logger.SetOutput(log.FILE, map[string]interface{}{
		"path":        l.GetPath(),
		"rotate":      l.GetRotate(),
		"daily":       l.GetDaily(),
		"compress":    l.GetCompress(),
		"maxlines":    l.GetMaxLines(),
		"maxsize":     l.GetMaxLines(),
		"maxkeepdays": l.GetMaxKeepDays(),
	})
}
