package main

import (
	"embed"
	"fmt"
	"github.com/zserge/lorca"
	"log"
	"main/binds"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
)

//go:embed www
var fs embed.FS
var ui lorca.UI
var ln net.Listener

func updateScene() {
	if ui != nil {
		ui.Done()
		ui.Close()
	}
	if ln != nil {
		ln.Close()
	}
	args := []string{}
	if runtime.GOOS == "linux" {
		args = append(args, "--class=Lorca")
	}
	var err error
	ui, err = lorca.New("", "", 1280, 720, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	binds.SetUp(ui)
	ln, err = net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(http.FS(fs)))
	ui.Load(fmt.Sprintf("http://%s/www", ln.Addr()))
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Println("exiting...")
}
