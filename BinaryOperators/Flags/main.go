package main

import "fmt"

// Flag ...
type Flag uint8

// Const ...
const (
	FIRST  Flag = 1 << iota // 1 << 0 = 0000 0001
	SECOND                  // 1 << 1 = 0000 0010
	THIRD
)

func foobar(flag Flag) {
	if (flag & FIRST) != 0 { // Check
		fmt.Print("FIRST is set: ")
	} else {
		fmt.Print("FIRST isn't set: ")
	}

	fmt.Printf("%08b \n", flag)
}

func main() {
	var f Flag

	f = f | FIRST // Set
	foobar(f)

	f = f | FIRST | SECOND // Set
	foobar(f)

	f = f &^ FIRST // Clear
	foobar(f)

	f = f ^ FIRST // Toggle
	foobar(f)

	foobar(SECOND)
	foobar(FIRST | THIRD | SECOND)
}
