package main

import (
	"fmt"
	"syscall/js"
)

var document js.Value
var body js.Value
var empty []js.Value

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
				createElement("h1", "", empty, "12"),
				createElement("h1", "", empty, "123"),
				createElement("h1", "", empty, "123"),
			},
			""),
	)
	<-make(chan bool)
}
