# Notebook

Just another notebook.

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

