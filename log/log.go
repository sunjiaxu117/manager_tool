package log

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var (
	MLog *logrus.Logger
)

func InitLogger() {
	MLog = logrus.New()
	MLog.SetFormatter(&logrus.JSONFormatter{})

	MLog.SetOutput(&lumberjack.Logger{
		// 日志输出文件路径。
		Filename: "log/logs/manager.log",
		// 日志文件最大 size, 单位是 MB。
		MaxSize: 1024,
		// 最大过期日志保留的个数。
		MaxBackups: 10,
		// 保留过期文件的最大时间间隔，单位天。
		MaxAge: 7,
	})
}
