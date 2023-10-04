package log

import "go.uber.org/zap"

var (
	Logger *zap.Logger
)

func Init() {
	Logger, _ = zap.NewDevelopment()
}

// Debug, Info, Warning, Error, DPanic, Panic, and Fatal.
