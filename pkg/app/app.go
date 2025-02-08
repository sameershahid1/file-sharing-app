package app

import (
	"context"
	"fmt"
	"sync"

	"github.com/libp2p/go-libp2p/core/host"
)

var CurrentMultiAddress string

// App struct
type App struct {
	ctx         context.Context
	HostAddress string
	Host        *host.Host
	mu          sync.Mutex
}

// NewApp creates a new App application struct
func NewApp(myHost *host.Host, hostAddress string) *App {
	return &App{
		HostAddress: hostAddress,
		Host:        myHost,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetMultiAddress() string {
	return a.HostAddress
}
