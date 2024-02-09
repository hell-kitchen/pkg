package confita

import (
	"github.com/heetch/confita"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_setNewLoader(t *testing.T) {
	loader := Get()
	otherLoader := confita.NewLoader()
	require.NotEqual(t, loader, otherLoader)
	assert.Equal(t, loader, _loader)
	restore := SetNewLoader(otherLoader)
	now := Get()
	assert.Equal(t, now, otherLoader)
	restore()
	now = Get()
	assert.NotEqual(t, now, otherLoader)
	assert.Equal(t, now, loader)
}
