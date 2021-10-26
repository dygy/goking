package main

import (
	"embed"
	"github.com/zserge/lorca"
	"net"
)

//go:embed www
var fs embed.FS
var ui lorca.UI
var ln net.Listener

func main() {
	updateScene()
}
