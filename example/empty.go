// Public domain 2014, author: Sonia Keys.

package main

import (
	"fmt"

	"github.com/soniakeys/exit"
)

func main() {
	defer exit.Handler()
	f()
	fmt.Println("f returned")
}

func f() {
	defer fmt.Println("f cleanup")
	fmt.Println("f attempt")
	if !false {
		exit.Log("")
	}
	fmt.Println("f completed")
}
