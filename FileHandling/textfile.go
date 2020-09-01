package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	content  []byte = make([]byte, 0)
	filename string = "example.text"
)

func readEntireFile() {
	var err error

	content, err = ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Print(string(content))
}

func writeEntireFile() {
	fmt.Print(string(content))

	/* If the file does not exist, WriteFile creates it
	 * with permissions (in this case 0644); otherwise
	 * WriteFile truncates it before writing, without
	 * changing permissions.
	 *
	 * https://golang.org/pkg/io/ioutil/#WriteFile
	 */

	err := ioutil.WriteFile(filename, content, 0644)

	if err != nil {
		panic(err)
	}
}

func readLineByLineText() {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func appendText() {
	/*
	 * !! Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	 *
	 * O_RDONLY  = open the file read-only.
	 * O_WRONLY  = open the file write-only.
	 * O_RDWR    = open the file read-write.
	 *
	 * !! The remaining values may be or'ed in to control behavior.
	 *
	 * O_APPEND   = append data to the file when writing.
	 * O_CREATE   = create a new file if none exists.
	 * O_EXCL     = used with O_CREATE, file must not exist.
	 * O_SYNC     = open for synchronous I/O.
	 * O_TRUNC    = truncate regular writable file when opened.
	 *
	 * https://golang.org/pkg/os/#pkg-constants
	 */

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	if _, err := file.Write([]byte("*** APPEND ***")); err != nil {
		panic(err)
	}

	file.Sync()

	fmt.Println("\nText is appended to the file.")
}

func main() {
	fmt.Println("(1) Func readEntireFile BEGIN")
	readEntireFile()
	fmt.Println("(2) Func readEntireFile END")

	fmt.Println("\n(3) Func readLineByLine BEGIN")
	readLineByLineText()
	fmt.Println("(4) Func readLineByLine END")

	appendText()

	fmt.Println("\n(5) Func readLineByLine BEGIN")
	readLineByLineText()
	fmt.Println("(6) Func readLineByLine END")

	fmt.Println("\n(7) Func writeEntireFile BEGIN")
	writeEntireFile()
	fmt.Println("(8) Func writeEntireFile END")
}
