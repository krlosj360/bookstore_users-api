package logger

import (
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/url"
)

const (
	// DPanic, Panic and Fatal level can not be set by user
	DebugLevelStr   string = "debug"
	InfoLevelStr    string = "info"
	WarningLevelStr string = "warning"
	ErrorLevelStr   string = "error"
)

var (
	globalLogger *zap.Logger
	devMode      bool = false
)

type lumberjackSink struct {
	*lumberjack.Logger
}

func (lumberjackSink) Sync() error {
	return nil
}

func init() {

	encoderConfig := zapcore.EncoderConfig{
		LevelKey:     "level",
		TimeKey:      "time",
		MessageKey:   "msg",
		EncodeTime:   zapcore.ISO8601TimeEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	ll := lumberjack.Logger{
		Filename:   "C:/Users/carlo/Documents/log/app.log",
		MaxSize:    1024, //MB
		MaxBackups: 30,
		MaxAge:     90, //days
		Compress:   true,
	}
	zap.RegisterSink("lumberjack", func(*url.URL) (zap.Sink, error) {
		return lumberjackSink{
			Logger: &ll,
		}, nil
	})
	loggerConfig := zap.Config{
		Level:         zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:      "json",
		EncoderConfig: encoderConfig,
		OutputPaths:   []string{fmt.Sprintf("lumberjack:%s", "C:/Users/carlo/Documents/log/app.log")},
	}

	var err error
	if globalLogger, err = loggerConfig.Build(); err != nil {
		panic(err)
	}
	/*globalLogger, err := loggerConfig.Build()
	if err != nil {
		panic(fmt.Sprintf("build zap logger from config error: %v", err))
	}
	//zap.ReplaceGlobals(globalLogger)

	*/

}

func Info(msg string, tags ...zap.Field) {
	globalLogger.Info(msg, tags...)
	globalLogger.Sync()
}

func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	globalLogger.Error(msg, tags...)
	globalLogger.Sync()
}
