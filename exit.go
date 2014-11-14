// Public domain 2014, author: Sonia Keys.

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
//     defer exit.Handler() // do this once at the beginning of the program
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
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

type fatal struct {
	code int
	msg  string
}

// Handler is the fatal error handler.  Defer at the start of main() before
// calling any of Code, Log, Logf, CodeLog, or CodeLogf.
//
// Handler handles paincs originating from these functions by logging messages
// with the standard logger of the log package and then exiting
// with an exit code.
//
// By default, output of the standard logger goes to stderr.
// If stderr is going to a terminal, Handler first sets logging flags to 0
// to suppress the date and time stamp.  In the case where stderr is
// redirected to a file for example, Handler does not reset the flags.
func Handler() {
	err := recover()
	if err == nil {
		return
	}
	f, ok := err.(fatal)
	if !ok {
		panic(err)
	}
	if f.msg > "\n" {
		if terminal.IsTerminal(int(os.Stderr.Fd())) {
			log.SetFlags(0)
		}
		log.Print(f.msg)
	}
	os.Exit(f.code)
}

// Code calls any deferred functions, then exits the program with the given
// exit code.
func Code(code int) {
	panic(fatal{1, ""})
}

// CodeLog calls any deferred functions, logs a message as with log.Println,
// then exits the program with the given exit code.
func CodeLog(code int, logArgs ...interface{}) {
	panic(fatal{code, fmt.Sprintln(logArgs...)})
}

// CodeLogf calls any deferred functions, logs a message as with log.Printf,
// then exits the program with the given exit code.
func CodeLogf(code int, logFormat string, logArgs ...interface{}) {
	panic(fatal{code, fmt.Sprintf(logFormat, logArgs...)})
}

// Log calls any deferred functions, logs a message as with log.Println,
// then exits the program with an exit code of 1.
func Log(args ...interface{}) {
	panic(fatal{1, fmt.Sprintln(args...)})
}

// Logf calls any deferred functions, logs a message as with log.Printf,
// then exits the program with an exit code of 1.
func Logf(format string, args ...interface{}) {
	panic(fatal{1, fmt.Sprintf(format, args...)})
}
