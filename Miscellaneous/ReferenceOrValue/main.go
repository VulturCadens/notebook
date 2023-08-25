package main

import "fmt"

type foo struct {
	x int
}

func fnArray(a [1]int) {
	a[0] = 42;	
	fmt.Printf("    fnArray: %d\n", a[0]);
}

func fnSlice(a []int) {
	a[0] = 42;
	fmt.Printf("    fnArray: %d\n", a[0]);	
}

func fnStruct(a foo) {
	a.x = 42;
	fmt.Printf("    fnStruct: %d\n", a.x);
}

func fnPointer(a *int) {
	var x int = 42;
	a = &x;
	fmt.Printf("    fnPointer: %d\n", *a);
}

func main() {
	fmt.Println("Struct");
	var s = foo{x: 1};
	fmt.Printf("    main: %d\n", s.x);	
	fnStruct(s);
	fmt.Printf("    main: %d\n\n", s.x);
	
	fmt.Println("Array");
	var v = [1]int{1};
	fmt.Printf("    main: %d\n", v[0]);
	fnArray(v);
	fmt.Printf("    main: %d\n\n", v[0]);
	
	fmt.Println("Slice");
	fmt.Printf("    main: %d\n", v[0]);	
	fnSlice(v[:]);
	fmt.Printf("    main: %d\n\n", v[0]);

	fmt.Println("Pointer");
	var x int = 1;
	var p *int = &x; // Produces the slice.
	fmt.Printf("    main: %d\n", *p);	
	fnPointer(p);
	fmt.Printf("    main: %d\n", *p);	
}
