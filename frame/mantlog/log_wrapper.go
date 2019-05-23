package mantlog

// 关闭日志文件句柄
func Close() {
	Logger.Close()
}

// 日志级别：Debug
func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

// 日志级别：Debugf
func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

// 日志级别：Info
func Info(args ...interface{}) {
	Logger.Info(args...)
}

// 日志级别：Infof
func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

// 日志级别：Warn
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

// 日志级别：Warnf
func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}

// 日志级别：Error
func Error(args ...interface{}) {
	Logger.Error(args...)
}

// 日志级别：Errorf
func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

// 日志级别：Fatal
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

// 日志级别：Fatalf
func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}
