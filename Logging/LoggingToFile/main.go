package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.Create("./logfile.log")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	/*
	 * Ldate          The date in the local time zone: 2009/01/23
	 * Ltime          The time in the local time zone: 01:23:23
	 * Lmicroseconds  Microsecond resolution: 01:23:23.123123.  assumes Ltime.
	 * Llongfile      Full file name and line number: /a/b/c/d.go:23
	 * Lshortfile     Final file name element and line number: d.go:23. overrides Llongfile
	 * LUTC           If Ldate or Ltime is set, use UTC rather than the local time zone
	 * Lmsgprefix     Move the "prefix" from the beginning of the line to before the message
	 * LstdFlags      Equals Ldate | Ltime (initial values for the standard logger)
	 *
	 * https://golang.org/pkg/log/#pkg-constants
	 */

	logger := log.New(file, "Application ", log.Lmicroseconds|log.Lshortfile)

	logger.Print("First line.")

	for i := 2; i < 5; i++ {
		time.Sleep(time.Millisecond * 200)
		logger.Printf("%d. line.", i)
	}
}
