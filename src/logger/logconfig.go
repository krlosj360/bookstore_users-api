package logger

import (
	"go.uber.org/zap/zapcore"
	"os"
	"strconv"
)

type envVals struct {
	filePath string
	stdout   bool
	level    zapcore.Level
}

//getEnv ログに関する環境変数を設定
func getEnv() (*envVals, error) {
	res := envVals{}
	res.filePath = os.Getenv("LOGGER_FILE_PATH")

	var err error
	res.stdout, err = strconv.ParseBool(os.Getenv("LOGGER_STDOUT"))
	if err != nil {
		res.stdout = true
	}

	level := os.Getenv("LOGGER_LEVEL")
	switch level {
	case "debug":
		res.level = zapcore.DebugLevel
	case "error":
		res.level = zapcore.ErrorLevel
	default:
		res.level = zapcore.InfoLevel
	}

	return &res, nil
}
