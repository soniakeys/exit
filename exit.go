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

// Exit specifies an exit code and logger for logging a final message.
//
// If Logger is nil, the log package default logger is used.
type Exit struct {
	ExitCode int
	*log.Logger
}

// Default holds Exit setting for the Log and Logf functions.
var Default = Exit{ExitCode: 1}

type msg struct {
	Exit
	string
}

// Handler is the fatal error handler.  Defer at the start of main() before
// calling any of Code, Log, Logf, Exit.Log, or Exit.Logf.
//
// When Handler is deferred first, Go ensures that all other deferred functions
// are run first.  Handler then handles paincs originating from Code, Log,
// Logf, Exit.Log, or Exit.Logf by logging a message and then exiting with
// an exit code.
//
// There are a couple of frills.  Empty messages are not logged; the program
// simply exits with the specified code.  With non-empty messages, if a logger
// is not specified and stderr seems to be directed to a terminal, Handler sets
// logging flags 0 to suppress the usual date and time stamp.
func Handler() {
	err := recover()
	if err == nil {
		return
	}
	m, ok := err.(msg)
	if !ok {
		panic(err)
	}
	if m.string > "\n" {
		switch {
		case m.Logger != nil:
			m.Logger.Print(m.string)
		case terminal.IsTerminal(int(os.Stderr.Fd())):
			log.SetFlags(0)
			fallthrough
		default:
			log.Print(m.string)
		}
	}
	os.Exit(m.ExitCode)
}

// Log runs any deferred functions, logs a message as with log.Println,
// then exits the program with an exit code.
//
// The logger and exit code used are those of exit.Default.
func Log(args ...interface{}) {
	panic(msg{Default, fmt.Sprintln(args...)})
}

// Logf calls any deferred functions, logs a message as with log.Printf,
// then exits the program with an exit code.
//
// The logger and exit code used are those of exit.Default.
func Logf(format string, args ...interface{}) {
	panic(msg{Default, fmt.Sprintf(format, args...)})
}

// Code runs any deferred functions, then exits the program with the given
// exit code.
func Code(code int) {
	Default.ExitCode = code
	panic(msg{Default, ""})
}

// Log runs any deferred functions, logs a message as with log.Println,
// then exits the program with an exit code.
//
// The logger and exit code used are those of the receiver x.
func (x Exit) Log(args ...interface{}) {
	panic(msg{x, fmt.Sprintln(args...)})
}

// Logf calls any deferred functions, logs a message as with log.Printf,
// then exits the program with an exit code.
//
// The logger and exit code used are those of the receiver x.
func (x Exit) Logf(format string, args ...interface{}) {
	panic(msg{x, fmt.Sprintf(format, args...)})
}
