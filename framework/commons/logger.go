package commons

import (
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func InitLogger() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
}

func Logger() *zap.Logger {
	return logger
}

func Error() {

}

func Info(msg string) {

}
