// Public domain 2014, author: Sonia Keys.

package main

import (
	"fmt"

	"github.com/soniakeys/exit"
)

func main() {
	defer exit.Handler() // do this once at the beginning of the program
	f()
	fmt.Println("f returned")
}

func f() {
	defer fmt.Println("f cleanup")
	fmt.Println("f attempt")
	if !false {
		exit.Log("Fatal error in f") // use exit.Log anywhere
	}
	fmt.Println("f completed")
}
