package logger

import "go.uber.org/zap"

func Initialize() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return logger
}
