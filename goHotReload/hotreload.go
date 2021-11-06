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
	"sync"
	"time"
)

var runCMD []*exec.Cmd
var stderr bytes.Buffer
var wg sync.WaitGroup

func main() {
	go func() {
		hotReload()
	}()
	for {
		fmt.Printf("%v+\n", time.Now())
		wg.Wait()
		runCMD = append(runCMD, exec.Command("go", "run", "main"))
		last().Dir, _ = filepath.Abs("")
		err := last().Run()
		last().Stderr = &stderr
		if err != nil {
			fmt.Println("starter " + fmt.Sprint(err))
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
				if event.Name() == "src" {
					rerun(true)
				} else if event.Name() == "binds" {
					rerun(false)
				}
			case err := <-w.Error:
				log.Println(err)
			case <-w.Closed:
				return
			}
		}
	}()

	addSource(w, "./binds")
	addSource(w, "./goking/src")

	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Println(err)
	}
}

func last() *exec.Cmd {
	if len(runCMD) > 0 {
		return runCMD[len(runCMD)-1]
	} else {
		return nil
	}
}

func addSource(w *watcher.Watcher, path string) {
	if err := w.AddRecursive(path); err != nil {
	}
}

func rerun(isReact bool) {
	wg.Add(1)
	log.Println("forceClearCache")
	log.Println("taskkill", "/T", "/F", "/PID", strconv.Itoa(last().Process.Pid))
	kill := exec.Command("taskkill", "/T", "/F", "/PID", strconv.Itoa(last().Process.Pid))
	kill.Stderr = &stderr
	err := kill.Run()
	if err != nil {
		fmt.Println("kill " + fmt.Sprint(err))
	}

	if isReact {
		rebuildCMD := exec.Command("npm", "run", "standard-build")
		dir, _ := filepath.Abs("")
		rebuildCMD.Dir = path.Join(dir, "/goking/")
		err := rebuildCMD.Run()
		if err != nil {
			log.Println(err)
		}
	}
	log.Println(wg)
	wg.Done()
}
