// +bulid wasm

package main

import (
	"fmt"
	"log"
	"syscall/js"
)

func main() {

	log.Printf("main() function here\n")

	fn := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Printf("got into the function\n")
		log.Printf("a log.Printf statement 1\n")
		return nil
	})
	js.Global().Get("window").Call("setCallback", fn)

}
