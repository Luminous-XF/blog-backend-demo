package util

import (
	"fmt"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger
)

func InitLogger(configFile string) {
	viper := ParseConfig(configFile)
	Logger = logrus.New()

	switch strings.ToLower(viper.GetString("level")) {
	case "debug":
		Logger.SetLevel(logrus.DebugLevel)
	case "info":
		Logger.SetLevel(logrus.InfoLevel)
	case "warn":
		Logger.SetLevel(logrus.WarnLevel)
	case "error":
		Logger.SetLevel(logrus.ErrorLevel)
	case "panic":
		Logger.SetLevel(logrus.PanicLevel)
	default:
		panic(fmt.Errorf("invalid log level: %s", viper.GetString("level")))
	}

	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05.000",
	})

	logFile := ProjectRootPath + viper.GetString("file")
	fout, err := rotatelogs.New(
		// 指定日志文件的路径和名称, 路径不存在时会创建
		logFile+".%Y%m%d%H",
		// 为最新的一份日志创建软链接
		rotatelogs.WithLinkName(logFile),
		// 每隔一小时生成一份新的日志文件
		rotatelogs.WithRotationTime(time.Hour*1),
		// 只保留近 7 天的日志, 或使用 WithRotationCount 只保留最近的几份日志
		rotatelogs.WithMaxAge(7*24*time.Hour),
	)

	if err != nil {
		panic(err)
	}

	// 设置日志文件
	Logger.SetOutput(fout)
	// 输出是从哪里调起的日志打印
	Logger.SetReportCaller(true)
}
