// Public domain 2014, author: Sonia Keys.

package main

import (
	"fmt"

	"github.com/soniakeys/exit"
)

func main() {
	go fish()
	select {}
}

func fish() {
	// Okay so the main example says do this once at the beginning of the
	// program but of course you really need to do it at the beginning of
	// any goroutine that might call one of Code, Log, or so on.
	defer exit.Handler()
	f()
	fmt.Println("f returned")
}

func f() {
	defer fmt.Println("f cleanup")
	fmt.Println("f attempt")
	if !false {
		// use exit.Log in any goroutine where you've defered a Hanlder.
		// (remember though, you're exiting the whole program, not just
		// this goroutine.)
		exit.Log("Fatal error in f")
	}
	fmt.Println("f completed")
}
