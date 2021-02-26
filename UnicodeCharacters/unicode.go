package main

import (
	"fmt"
	"unicode/utf8"
)

var unicodeString string = "aωA◆"

func main() {
	/*
	 * Rune (type Rune int32) holds any unicode character.
	 */

	fmt.Printf("There is %d runes in string (%s). \n", utf8.RuneCountInString(unicodeString), unicodeString)
	fmt.Printf("Length of string is %d bytes. \n\n", len(unicodeString))

	fmt.Println("Iterate over the characters of a string (the range expression).")
	for i, r := range unicodeString {
		fmt.Printf("  Index: %d  Character: %c  Unicode: %d\n", i, r, r)
	}

	fmt.Println("\nIterate over the bytes of a string.")
	for i, b := range []byte(unicodeString) {
		fmt.Printf("  Index: %-2d  Byte: %d\n", i, b)
	}

	fmt.Println("\nIterate over the characters of a string (utf8.DecodeRuneInString).")
	for i := 0; i < len(unicodeString); {
		rune, size := utf8.DecodeRuneInString(unicodeString[i:])
		fmt.Printf("  Index: %d  Size (bytes): %d  Character: %c\n", i, size, rune)
		i += size
	}
}
