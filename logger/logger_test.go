package logger

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"testing"
)

func TestReplaces(t *testing.T) {
	before := zap.L()
	logger, err := NewProduction()
	require.NoError(t, err)
	if assert.NotNil(t, logger) {
		after := zap.L()
		assert.Equal(t, after, logger)
		assert.NotEqual(t, after, before)
	}
}
