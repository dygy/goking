package main

import (
	"embed"
	"github.com/zserge/lorca"
	"net"
	"runtime"
	"sync"
)

//go:embed www
var fs embed.FS
var ui lorca.UI
var ln net.Listener

// Go types that are bound to the UI must be thread-safe, because each binding
// is executed in its own goroutine. In this simple case we may use atomic
// operations, but for more complex cases one should use proper synchronization.
type counter struct {
	sync.Mutex
	count int
}

func (Controller *counter) getOS() string {
	return runtime.GOOS
}

func main() {
	updateScene()
}
