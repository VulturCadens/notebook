# Notebook

Just another notebook.

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
