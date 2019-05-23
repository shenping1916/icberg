package mantlog

func LevelDebug() string {
	return Logger.LevelString()[0]
}

func LevelInfo() string {
	return Logger.LevelString()[1]
}

func LevelWarn() string {
	return Logger.LevelString()[2]
}

func LevelError() string {
	return Logger.LevelString()[3]
}

func LevelFatal() string {
	return Logger.LevelString()[4]
}
