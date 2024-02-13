package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func highPriorityLevelEnablerFunc(lvl zapcore.Level) bool {
	return lvl >= zapcore.ErrorLevel
}

func lowPriorityLevelEnablerFunc(lvl zapcore.Level) bool {
	return lvl < zapcore.ErrorLevel
}

// NewProduction creates production logger with provided options.
func NewProduction(opts ...zap.Option) (*zap.Logger, error) {
	highPriority := zap.LevelEnablerFunc(highPriorityLevelEnablerFunc)
	lowPriority := zap.LevelEnablerFunc(lowPriorityLevelEnablerFunc)

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleErrors := zapcore.Lock(os.Stderr)
	jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, consoleErrors, highPriority),
		zapcore.NewCore(jsonEncoder, consoleDebugging, lowPriority),
	)
	logger := zap.New(core, opts...)
	zap.ReplaceGlobals(logger)
	return logger, nil
}
