package configuration

import "github.com/sirupsen/logrus"

func SetLogLevel(level Logs) logrus.Level {
	switch level.LogLevel {
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "trace":
		return logrus.TraceLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}
