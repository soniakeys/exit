// Public domain 2014, author: Sonia Keys.

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/soniakeys/exit"
)

var doge = exit.Exit{255, log.New(os.Stdout, "Wow ", log.Lshortfile)}

func main() {
	defer exit.Handler()
	f()
	fmt.Println("f returned")
}

func f() {
	defer fmt.Println("f cleanup")
	fmt.Println("f attempt")
	if !false {
		doge.Log("so scare")
	}
	fmt.Println("f completed")
}
