package main

import (
	"bytes"
	"fmt"
	"github.com/radovskyb/watcher"
	"log"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

var runCMD []*exec.Cmd
var stderr bytes.Buffer

func main() {
	for {
		fmt.Printf("%v+\n", time.Now())
		go func() { hotReload() }()
		runCMD = append(runCMD, exec.Command("go", "run", "main"))
		last().Dir, _ = filepath.Abs("")
		err := last().Run()
		last().Stderr = &stderr
		if err != nil {
			fmt.Println("starter " + fmt.Sprint(err) + ": " + stderr.String())
		}
	}
}

func hotReload() {
	w := watcher.New()
	w.IgnoreHiddenFiles(true)

	go func() {
		for {
			select {
			case event := <-w.Event:
				// Print the event's info.
				log.Println(event.Name())
				if event.Name() == "src" {
					rebuildCMD := exec.Command("npm", "run", "standard-build")
					dir, _ := filepath.Abs("")
					rebuildCMD.Dir = path.Join(dir, "/goking/")
					out, err := rebuildCMD.Output()
					if err != nil {
						log.Println(err)
					}
					if out != nil {
						log.Println("forceClearCache")
						log.Println("taskkill", "/T", "/F", "/PID", strconv.Itoa(last().Process.Pid))
						kill := exec.Command("taskkill", "/T", "/F", "/PID", strconv.Itoa(last().Process.Pid))
						kill.Stderr = &stderr
						err := kill.Run()
						if err != nil {
							fmt.Println("kill " + fmt.Sprint(err) + ": " + stderr.String())
						}
					}
				}
			case err := <-w.Error:
				log.Println("exit?2")
				log.Println(err)
			case <-w.Closed:
				log.Println("exit?1")
				hotReload()
				return
			}
		}
	}()
	if err := w.AddRecursive("./goking/src"); err != nil {
		log.Println(err)
		log.Println("exit?5")
	}

	// Start the watching process - it'll check for changes every 100ms.
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Println(err)
		log.Println("exit?4")
	}
	log.Println("exit?3")
}

func last() *exec.Cmd {
	return runCMD[len(runCMD)-1]
}
