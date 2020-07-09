package common

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"path"
	"time"
)

var Log = logrus.New()

func InitLog(){

	now := time.Now()
	logFilePath := viper.GetString("LOG_FILE_PATH")
	if logFilePath == ""{
		if dir, err := os.Getwd(); err == nil {
			logFilePath = dir + "/logs/"
		}
		if err := os.MkdirAll(logFilePath, 0777); err != nil {
			log.Printf(err.Error())
		}
	}

	logFileName := viper.GetString("LOG_FILE_NAME")

	if logFileName == ""{
		logFileName = now.Format("2006-01-02") + ".log"
	}
	//日志文件
	fileName := path.Join(logFilePath, logFileName)

	//写入文件
	var writer io.Writer
	logFileType := viper.GetString("LOG_FILE_TYPE")
	if logFileType == "" || logFileType == "Stdout"{
		writer = os.Stdout
	} else {
		var err error
		writer, err = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			log.Panic("create file log.txt failed: %v", err)
		}
	}

	Log.SetOutput(io.MultiWriter(writer))
	Log.SetLevel(logrus.InfoLevel)
	Log.SetFormatter(&logrus.TextFormatter{})
}

func Logger() *logrus.Logger{
	return Log
}

func LoggerToFile() gin.HandlerFunc {
	//设置输出
	logger := Logger()
	logger.SetFormatter(&logrus.TextFormatter{})
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		//日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

