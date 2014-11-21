// Public domain 2014, author: Sonia Keys.

package main

import (
	"fmt"
	"log"

	"github.com/soniakeys/exit"
)

func main() {
	defer exit.Handler()
	exit.Default.ExitCode = 255
	log.SetPrefix("Wow ")
	f()
	fmt.Println("f returned")
}

func f() {
	defer fmt.Println("f cleanup")
	fmt.Println("f attempt")
	if !false {
		exit.Log("so scare")
	}
	fmt.Println("f completed")
}
