package main

import (
	"fmt"
	"os"
)

var (
	file *os.File
	err  error

	filename string = "example.binary"

	buffer []byte  = make([]byte, 6)
	array  [2]byte = [2]byte{0x11, 0x22}
)

func main() {
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

	if file, err = os.OpenFile(filename, os.O_SYNC|os.O_RDWR, 0644); err != nil {
		panic(err)
	}

	defer file.Close()

	// Read(b []byte) (n int, err error)
	if _, err = file.Read(buffer); err != nil {
		panic(err)
	}

	for i := range buffer {
		// \x20 = whitespace symbol U+0020.
		fmt.Printf("%d. byte is 0x%X (%8b) \x20", i, buffer[i], buffer[i])
	}
	fmt.Println()

	// WriteAt(b []byte, off int64) (n int, err error)
	if _, err = file.WriteAt(array[:], 2); err != nil {
		panic(err)
	}

	// Should use file.Sync(), if asynchronous I/O.

	/* Seek(offset int64, whence int) (ret int64, err error)
	 *
	 * whence: 0 means relative to the origin of the file.
	 *         1 means relative to the current offset.
	 *         2 means relative to the end.
	 *
	 * https: //golang.org/pkg/os/#File.Seek
	 */

	// Set the offset for the next Read().
	if _, err = file.Seek(0, 0); err != nil {
		panic(err)
	}

	if _, err = file.Read(buffer); err != nil {
		panic(err)
	}

	for i := range buffer {
		fmt.Printf("%d. byte is 0x%X (%08b) \x20", i, buffer[i], buffer[i])
	}
	fmt.Println()

	// Restore an example file.
	if _, err = file.WriteAt([]byte{0xCC, 0xDD}, 2); err != nil {
		panic(err)
	}
}
