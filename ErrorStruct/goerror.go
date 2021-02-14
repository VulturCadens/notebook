package main

import (
	"fmt"
)

/*
 *  The error type is an interface type.
 *
 *  type error interface {
 *      Error() string
 *  }
 */

type xError struct {
	name string
	code int
}

func (e xError) Error() string {
	return fmt.Sprintf("%s (%d)", e.name, e.code)
}

func getError() error {
	return &xError{
		"Not Found",
		404,
	}
}

func main() {
	if err := getError(); err != nil {

		x := err.(*xError)

		fmt.Println(err) // ok

		fmt.Printf("\"%s\" response code is %d.\n", x.name, x.code) // ok

		/*
			fmt.Printf("\"%s\" response code is %d.\n", err.name, err.code) // err.name undefined (type error has no field or method name)
		*/
	}
}
