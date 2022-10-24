package logger

import (
	"bullbeckt/setting"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var lg *zap.Logger

func Init(cfg *setting.LogConfig, mode string) (err error) {

}

// 构造写日志结构体
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumbLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		MaxSize:    maxSize,
	}
	return zapcore.AddSync(lumbLogger)
}

// aop 也就是gin的中间件，请求进入前打印日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL

	}
}
