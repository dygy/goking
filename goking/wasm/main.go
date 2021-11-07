package main

import (
	"fmt"
	"runtime"
	"syscall/js"
)

var document js.Value
var body js.Value
var empty []js.Value

func Hello(this js.Value, args []js.Value) interface{} {
	message := args[0].String() // get the parameters
	return "Hello " + runtime.GOOS + message
}

func main() {
	fmt.Println("👋 Hello World! 🌍")
	document = js.Global().Get("document")
	body = document.Get("body")

	js.Global().Set("Hello", js.FuncOf(Hello))
	appendChild(
		createElement(
			"div",
			`
					background-color: green;
					display: flex;
					justify-content: space-around;
				`,
			[]js.Value{
				createElement("h1", "", empty, "123"),
				createElement("h1", "", empty, "123"),
				createElement("h1", "", empty, "123"),
			},
			""),
	)
	<-make(chan bool)
}

func createElement(tag string, styles string, innerContent []js.Value, innerText string) js.Value {
	element := document.Call("createElement", tag)
	if len(styles) > 0 {
		element.Set("style", styles)
	}
	for i := 0; i < len(innerContent); i++ {
		element.Call("appendChild", innerContent[i])
	}
	if len(innerText) > 0 {
		element.Set("innerText", innerText)

	}
	return element
}

func appendChild(child js.Value) {
	body.Call("appendChild", child)
}
