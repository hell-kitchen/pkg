package confita

import (
	"context"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/flags"
	"sync"
)

type Loader interface {
	// Load analyses all the Fields of the given struct for a "config" tag and queries each backend
	// in order for the corresponding key. The given context can be used for timeout and cancelation.
	Load(ctx context.Context, to interface{}) error
}

var (
	_loader Loader
	_mu     sync.Mutex
	_once   sync.Once
)

func initLoader() {
	_mu.Lock()
	_loader = confita.NewLoader(
		env.NewBackend(),
		flags.NewBackend(),
	)
	_mu.Unlock()
}

// Get return global loader.
func Get() (l Loader) {
	_once.Do(initLoader)
	_mu.Lock()
	l = _loader
	_mu.Unlock()

	return l
}

// SetNewLoader sets loader l as global.
//
// Return function on calling it will set old global loader back.
func SetNewLoader(l Loader) func() {
	_mu.Lock()
	old := _loader
	_loader = l
	_mu.Unlock()

	return func() {
		SetNewLoader(old)
	}
}
