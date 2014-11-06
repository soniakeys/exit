// Public domain 2014.

// Exit provides a way to exit from a program that runs deferred functions,
// logs a message to stderr, and then exits with an operating system exit code.
//
// The following example program is in source subdirectory "example".
//
//   package main
//
//   import (
//      "fmt"
//
//      "github.com/soniakeys/exit"
//   )
//
//   func main() {
//     defer exit.Exit() // do this once at the beginning of the program
//     f()
//     fmt.Println("f returned")
//   }
//
//   func f() {
//      defer fmt.Println("f cleanup")
//      fmt.Println("f attempt")
//      if !false {
//         exit.Log("Fatal error in f") // use exit.Log anywhere
//      }
//      fmt.Println("f completed")
//   }
//
// This prints "f attempt" and "f cleanup" to stdout, logs "Fatal error in f"
// to stderr, and exits with a program exit code of 1.
package exit

import (
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
)

type fatal struct {
	err interface{}
}

// Exit is the fatal error handler.  Defer at the start of main().
func Exit() {
	if err := recover(); err != nil {
		if f, ok := err.(fatal); ok {
			if terminal.IsTerminal(int(os.Stderr.Fd())) {
				log.SetFlags(0)
			}
			log.Fatal(f.err)
		}
		panic(err)
	}
}

// Log exits the program, first calling deferred functions, then calling
// log.Fatal(err), then exiting with an exit code of 1.
//
// If stderr is going to the terminal, logging flags are set to 0 to suppress
// the date and time stamp.  When stderr is redirected, err is logged with
// the standard logging flags.
func Log(err interface{}) {
	panic(fatal{err})
}
