##Exit

Exit provides a way to exit from a program while both running deferred
functions and returning a non-zero exit code.

Deferred functions are great for resource clean up.  os.Exit exits a program
immediately with an exit code.  This package provides a simple way to exit a
program immediately with an exit code, but also run the deferred functions
on the way out.

Of course it's just an application of panic and recover but the package saves
you writing that little bit of code for every command line program where you
want this capability.

###Status

The package is totally useful just as it is, but really it's more of a
technique demonstration.  It's missing features you might want, like exiting
with a given exit code instead of always just 1.  It also has a fluff feature,
suppressing the time stamp on stderr to the terminal.

Anyway, it's public domain, so copy and modify it according to your needs.

###Example

```
package main

import (
   "fmt"

   "github.com/soniakeys/exit"
)

func main() {
  defer exit.Exit() // do this once at the beginning of the program
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
$ go run example.go
f attempt
f cleanup
Fatal error in f
exit status 1
```

```
$ go build example.go
$ ./example.go 2>log
f attempt
f cleanup

$ echo $?
1

$ cat log
2014/10/02 20:03:59 Fatal error in f
```
