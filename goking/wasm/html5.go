package main

import "syscall/js"

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
