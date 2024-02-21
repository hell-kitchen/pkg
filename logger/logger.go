package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

const (
	ServiceField    = "service"
	InstanceIDField = "instance"
	LayerField      = "layer"
)

type Config struct {
	Service    *string
	InstanceID *string
}

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

func WithLayer(layer string) zap.Field {
	return zap.String(LayerField, layer)
}

func WithInstanceID(id string) zap.Field {
	return zap.String(InstanceIDField, id)
}

func WithService(service string) zap.Field {
	return zap.String(ServiceField, service)
}

func NewWithConfig(cfg Config, opts ...zap.Option) (*zap.Logger, error) {
	var fields []zap.Field
	if cfg.InstanceID != nil {
		fields = append(fields, WithInstanceID(*cfg.InstanceID))
	}
	if cfg.Service != nil {
		fields = append(fields, WithService(*cfg.Service))
	}

	opts = append(opts, zap.Fields(fields...))

	logger, err := NewProduction(opts...)

	zap.L().With()
	return logger, err
}
