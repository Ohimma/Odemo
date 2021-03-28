package middleware

import (
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ohimma/odemo/gin/config"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	log := logrus.New()
	// 生产环境写入文件JSON格式
	if config.Conf.Server.Mode == "release" {
		fileName := path.Join("./logs", "odemo.log")
		_, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		// logger.SetOutput(bufio.NewWriter(file))
		// writeFile(fileName, config.MaxAge)

		if err != nil {
			panic(err)
		}
		//设置日志格式、级别、输出地方
		log.SetFormatter(&logrus.JSONFormatter{})
		log.SetLevel(logrus.InfoLevel)

	} else {
		log.SetFormatter(&logrus.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
		log.WithFields(logrus.Fields{
			"app": "odemo",
		})
		log.SetLevel(logrus.DebugLevel)
		log.SetOutput(os.Stdout)
	}

	Logger = log
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		if uri == "/favicon.ico" {
			return
		}
		//开始时间
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		method := c.Request.Method
		statusCode := c.Writer.Status()
		ip := c.ClientIP()

		if config.Conf.Server.Mode == "release" {
			Logger.WithFields(logrus.Fields{
				"clientIp":    ip,
				"statusCode":  statusCode,
				"reqMethod":   method,
				"reqUri":      uri,
				"latencyTime": latencyTime,
			}).Info()
		} else {
			// now := time.Now().Format("2006-01-02 15:04:05")
			// Logger.Infof(" %3d | %13v | %15s | %s  %s",
			// 	statusCode,
			// 	latencyTime,
			// 	ip,
			// 	method,
			// 	uri,
			// )
			Logger.WithFields(logrus.Fields{
				"clientIp":    ip,
				"statusCode":  statusCode,
				"reqMethod":   method,
				"reqUri":      uri,
				"latencyTime": latencyTime,
			})
		}
	}
}

// 文件切割
// func writeFile(fileName string, maxAge time.Duration) {
// 	// 设置 rotatelogs
// 	logWriter, _ := rotatelogs.New(
// 		// 分割后的文件名称
// 		strings.Replace(fileName, ".log", "", -1)+".%Y-%m-%d.log",
// 		// 生成软链，指向最新日志文件
// 		rotatelogs.WithLinkName(fileName),
// 		// 设置最大保存时间
// 		rotatelogs.WithMaxAge(maxAge*24*time.Hour),
// 		// 设置日志切割时间间隔(1天)
// 		rotatelogs.WithRotationTime(24*time.Hour),
// 	)

// 	writeMap := lfshook.WriterMap{
// 		logrus.InfoLevel:  logWriter,
// 		logrus.FatalLevel: logWriter,
// 		logrus.DebugLevel: logWriter,
// 		logrus.WarnLevel:  logWriter,
// 		logrus.ErrorLevel: logWriter,
// 		logrus.PanicLevel: logWriter,
// 	}

// 	logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
// 		TimestampFormat: TIME_FORMAT,
// 	}))
// }

// 级别
// log.Trace("Something very low level.")
// log.Debug("Useful debugging information.")
// log.Info("Something noteworthy happened!")
// log.Warn("You should probably take a look at this.")
// log.Error("Something failed but I'm not quitting.")

// // Calls os.Exit(1) after logging
// log.Fatal("Bye.")
// // Calls panic() after logging
// log.Panic("I'm bailing.")
