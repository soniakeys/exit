// Public domain 2014, author: Sonia Keys.

package main

import (
	"fmt"
	"runtime"

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
		exit.Logf("Fatal error in f (%s)", runtime.Version())
	}
	fmt.Println("f completed")
}
