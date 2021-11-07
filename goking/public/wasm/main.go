package main

import (
	"fmt"
	"syscall/js"
)

func Hello(this js.Value, args []js.Value) interface{} {
	message := args[0].String() // get the parameters
	return "Hello " + message
}

func main() {
	fmt.Println("👋 Hello World! 🌍")
	js.Global().Set("Hello", js.FuncOf(Hello))

	<-make(chan bool)
}
