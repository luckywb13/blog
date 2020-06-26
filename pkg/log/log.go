package log

import (
	"github.com/luckywb13/blog/conf"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	DebugLevel int8 = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
)


var levelMap = map[string]int8{
	"debug": DebugLevel,
	"info":  InfoLevel,
	"warn":  WarnLevel,
	"error": ErrorLevel,
}


var logClient *zap.SugaredLogger

func InitLog(c *conf.Log) {

	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  c.Filename,
		MaxSize:   c.MaxSize,
		LocalTime: c.LocalTime,
		Compress:  c.Compress,
	})

	level := DebugLevel
	if tmpLevel, ok := levelMap[c.Level]; ok {
		level = tmpLevel
	}

	encoder := zap.NewProductionEncoderConfig()
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(zapcore.Level(level)))
	logger := zap.New(core)
	logClient = logger.Sugar()
}



func  Debug(v ...interface{}) {
	logClient.Debug(v...)
}

func Error(v ...interface{}) {
	logClient.Error(v...)
}

func Info(v ...interface{}) {
	logClient.Info(v...)
}

func Warn(v ...interface{}) {
	logClient.Warn(v...)
}

func DebugF(format string, v ...interface{}) {
	logClient.Debugf(format, v...)
}

func  ErrorF(format string, v ...interface{}) {
	logClient.Errorf(format, v...)
}

func InfoF(format string, v ...interface{}) {
	logClient.Infof(format, v...)
}

func WarnF(format string, v ...interface{}) {
	logClient.Warnf(format, v...)
}

func DebugW(msg string, v ...interface{}) {
	logClient.Debugw(msg, v...)
}

func  ErrorW(msg string, v ...interface{}) {
	 logClient.Errorw(msg, v...)
}

func InfoW(msg string, v ...interface{}) {
	logClient.Infow(msg, v...)
}

func WarnW(msg string, v ...interface{}) {
	logClient.Warnw(msg, v...)
}

func  DestroyLogger() {
	_ = logClient.Sync()
}
