# Exit

Exit provides a way to exit from a program while both running deferred
functions and returning a non-zero exit code.

[![GoDoc](https://godoc.org/github.com/soniakeys/exit?status.svg)](https://godoc.org/github.com/soniakeys/exit) [![status](https://sourcegraph.com/api/repos/github.com/soniakeys/exit/.badges/status.png)](https://sourcegraph.com/github.com/soniakeys/exit) [![Go Walker](http://gowalker.org/api/v1/badge)](https://gowalker.org/github.com/soniakeys/exit) [![GoSearch](http://go-search.org/badge?id=github.com%2Fsoniakeys%2Fexit)](http://go-search.org/view?id=github.com%2Fsoniakeys%2Fexit)

Deferred functions are great for resource clean up.  os.Exit exits a program
immediately with an exit code, but doesn't run your deferred functions.
This package provides a simple way to do both, exit a program immediately
with an exit code, but also run the deferred functions on the way out.

Of course it's just an application of panic and recover but the package saves
you writing that little bit of code for every command line program where you
want this capability.

## Example

```
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
```

```
$ go build example.go
$ ./example 2>log
f attempt
f cleanup

$ echo $?
1

$ cat log
2014/10/02 20:03:59 Fatal error in f
```

## See also

Similar package at https://github.com/youtube/vitess/tree/master/go/exit

