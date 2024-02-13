package logger

import (
	"fmt"
	"go.uber.org/zap"
)

// NewProduction creates production logger with provided options.
func NewProduction(opts ...zap.Option) (*zap.Logger, error) {
	logger, err := zap.NewProduction(opts...)
	if err != nil {
		return nil, fmt.Errorf("zap: NewProduction: %w", err)
	}
	return logger, nil
}
