package binds

import (
	"github.com/zserge/lorca"
	"log"
	"runtime"
	"sync"
)

type counter struct {
	sync.Mutex
	count int
}

func (Controller *counter) getOS() string {
	return runtime.GOOS
}

func (Controller *counter) getStr() string {
	return "LOL KEK"
}

func SetUp(ui lorca.UI) {
	ui.Bind("start", func() {
		log.Println("UI is on")
	})
	c := &counter{}
	ui.Bind("getOS", c.getOS)
	ui.Bind("getStr", c.getStr)
}
