# Notebook

Just another notebook.

## Restarting Server

"Command __entr__ - Run arbitrary commands when files change."

```bash
ls *.go|entr -r go run server.go
```

## The Go regexp package

* Supports UTF-8 encoded strings.

* Implements regular expressions with RE2 syntax.

* The regular expression syntax accepted by RE2: https://github.com/google/re2/wiki/Syntax

## Go Modules

* Reference: https://golang.org/ref/mod

* Tutorial: https://golang.org/doc/tutorial/create-module

## How patterns are matched

* From the most specific to the least specific.

* Handler registered for __/css__ works only for __/css__.

* The trailing slash __/css/__ will be matched for __/css/__, __/css/style.css__, __/css/font/default.woff__ and __/css__ with a redirect. ServeMux redirects a request for __/css__ to __/css/__ (HTTP status code 301).

* The slash __/__ matches all paths not matched by other patterns.

```go
http.HandleFunc("/", funcA)      /* matched for '/' '/foobar'      */
http.HandleFunc("/foo", funcB)   /* matched for '/foo'             */
http.HandleFunc("/foo/", funcC)  /* matched for '/foo/' '/foo/bar' */
```

## Types

### **Basic Types**

* Numbers (integer, float and complex), strings and booleans.

### **Composite Types**

* Arrays and structs.

```go
a := [3]int{0, 1, 2}

type s struct {
    foo string
    bar int
}
```

### **Composition**

```go
package main

import "fmt"

type common struct {
	name string
	age  int
}

type cat struct {
	common
	tailLength float32
}

type parrot struct {
	common
	wingLength float32
}

func main() {
	c := cat{common{"Garfield", 8}, 13.8}

	/*
	 *   The struct literal must follow the shape of the type declaration.
	 *   p := parrot{name: "Polly", age: 23, wingLength: 18.5} // Compile error!
	 */

	p := parrot{}
	p.name = "Polly"
	p.age = 23
	p.wingLength = 18.5

	fmt.Println(c)
	fmt.Println(p)
}
```

### **Reference Types**

* Pointers, slices, maps and channels.

```go
var p *int
var n int = 1
p = &n

s := []int{0, 1, 2}

m := make(map[string]int)     // m := map[string]int{}

c := make(chan int)


var x []byte              // Doesn't allocate memory and x -> nil.
var x = make([]byte, 0)   // Allocates memory and x -> memory.
```

`Array[:]` produces the slice of the underlying array.

```go
a := [2]int{1, 2}   // Array
s := a[:]           // Slice
```

The index of the element.

```go
var array = [...]int{10: 256, 5: 128} // Index 10 -> Value 256, Index 5 -> Value 128
fmt.Printf("%#v\n", array)
```

This is the output.
```bash
[11]int{0, 0, 0, 0, 0, 128, 0, 0, 0, 0, 256}
```

### **Interface types**

* Interface defines behaviours.

```go
type example interface {
    foo([]byte) string, error
    bar() int
}
```

## The Range Expression

* Array
* Pointer to an array
* Slice
* String
* Map
* Channel (receive operation)

```go
var array *[4]int = &[4]int{1, 2, 3, 4}

for index, value := range array {
	fmt.Printf("%d -> %d \n", index, value)
}
```

## Three dots

```go
package main

import (
	"fmt"
)

func foobar(args ...string) {     // Type is []string (slice).
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func main() {
	array := [...]string{"John", "Joan", "Jack"} // Type is [3]string (array).

	slice := array[:]     // Convert array to slice.

	foobar(slice...)      // Must be a slice.
}
```

## Function closure

It's ok to leave the channel open, and it's idiomatic Go. If a channel is unreachable, it will be garbage collected.

```go
package main

import (
        "fmt"
        "math/rand"
        "time"
)

func run() (chan int, func()) {
        var stop bool // The zero value is false for the boolean type.

        c := make(chan int)

        go func() {
                for !stop {
                        c <- rand.Intn(10)
                }
        }()

        return c, func() { stop = true }
}

func main() {
        rand.Seed(time.Now().UnixNano())

        c, stop := run()

        var r int

        for {
                r = <-c
                fmt.Println(r)
                if r == 8 {
                        stop()
                        break
                }
        }
}
```

## Type Conversion

The expression T(v) converts the value v to the type T.

```go
var str string = "Cat"

slice := []byte(str)
```

## Type assertion

The expression x.(T).

```go
var x interface{} = "Cat"   // The dynamic type of x is string.

str := x.(string)
```

A minimal working example.

```go
package main

import "fmt"

type receiver interface {
	receive()
}

type sender interface {
	send()
}

type general struct{}

func (g general) receive() {
	fmt.Println("Receive")
}

func (g general) send() {
	fmt.Println("Send")
}

func main() {
	var x receiver = general{}

	x.receive()

	if x, ok := x.(sender); ok {
		x.send()
	}
}
```

"The Flusher interface is implemented by ResponseWriters that allow an HTTP handler to flush buffered data to the client."

https://golang.org/pkg/net/http/#Flusher

```go
func (w http.ResponseWriter, req *http.Request) {
    f, ok := w.(http.Flusher)

    if !ok {
		http.Error(w, "500", http.StatusInternalServerError)
		return
    }
    
    for {
        // ...

        f.Flush() 
    }
}
```

## Context

"Package context defines the Context type, which carries deadlines, cancellation signals, and other request-scoped values across API boundaries and between processes."

https://pkg.go.dev/context

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func sub(ctx context.Context) {
    fmt.Print("Waiting cancelation... ")
    
    <-ctx.Done()

    fmt.Println("Done")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	sub(ctx)
}
```

## Channels

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan struct{})

	go func() {
		time.Sleep(2 * time.Second)
		channel <- struct{}{}
	}()

	<-channel // Blocks until the data is available in the channel.

	fmt.Println("OK")
}
```

Using a range loop to iterate over the values sent through the channel.

```go
package main

import (
	"fmt"
	"time"
)

var channel chan int

func send() {
	for x := 1; x < 10; x++ {
		time.Sleep(250 * time.Millisecond)
		channel <- x
	}

	close(channel)
}

func main() {
	channel = make(chan int)

	go send()

	for r := range channel {
		fmt.Println(r)
	}

	fmt.Println("The channel has been closed.")
}
```
