package main

import (
	"runtime"
	"syscall/js"
)

func Hello(this js.Value, args []js.Value) interface{} {
	message := args[0].String() // get the parameters
	return "Hello " + runtime.GOOS + message
}
