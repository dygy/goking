package main

import (
	"embed"
	"fmt"
	"github.com/radovskyb/watcher"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"sync"
	"time"

	"github.com/zserge/lorca"
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

func (Controller *counter) Add(n int) {
	Controller.Lock()
	defer Controller.Unlock()
	Controller.count = Controller.count + n
}

func (Controller *counter) Value() int {
	Controller.Lock()
	defer Controller.Unlock()
	return Controller.count
}
func (Controller *counter) getOS() string {
	return runtime.GOOS
}

func main() {
	updateScene(true)
}

func updateScene(isFirst bool) {
	if ui != nil && ln != nil {
		ui.Done()
		ui.Close()
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

	// A simple way to know when UI is ready (uses body.onload event in JS)
	ui.Bind("start", func() {
		log.Println("UI is ready")
	})
	// Create and bind Go object to the UI
	c := &counter{}
	ui.Bind("counterAdd", c.Add)
	ui.Bind("counterValue", c.Value)
	ui.Bind("getOS", c.getOS)

	// Load HTML.
	// You may also use `data:text/html,<base64>` approach to load initial HTML,
	// e.g: ui.Load("data:text/html," + url.PathEscape(html))

	ln, err = net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	go http.Serve(ln, http.FileServer(http.FS(fs)))
	ui.Load(fmt.Sprintf("http://%s/www", ln.Addr()))

	// You may use console.log to debug your JS code, it will be printed via
	// log.Println(). Also exceptions are printed in a similar manner.
	ui.Eval(`
		console.log("Hello, world!");
	`)
	if isFirst {
		hotload()
	}

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	log.Println("exiting...")
}

func hotload() {
	w := watcher.New()
	log.Println(111)

	// Ignore hidden files.
	w.IgnoreHiddenFiles(true)

	go func() {
		for {
			select {
			case event := <-w.Event:
				// Print the event's info.
				fmt.Println(event)
				path, err := os.Getwd()
				if err != nil {
					log.Println(err)
				}
				rebuildCMD := exec.Command("npm", "run", "standard-build")
				rebuildCMD.Dir = path + "\\goking\\"
				out, _ := rebuildCMD.Output()
				err = rebuildCMD.Start()
				if err != nil {
					log.Println(err)
				}
				if out != nil {
					log.Println(string(out))
					updateScene(false)
				}

				log.Println(rebuildCMD)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch test_folder recursively for changes.
	//
	// Watcher won't add .dotfile to the watchlist.
	if err := w.AddRecursive("./goking/src"); err != nil {
		log.Fatalln(err)
	}

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
