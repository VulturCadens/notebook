package main

import (
	"fmt"
	"unicode/utf8"
)

var unicodeString string = "ⓈωA◆‰"

func main() {
	/*
	 * Rune (type Rune int32) holds any unicode character.
	 */

	fmt.Printf("There is %d runes in string (%s). \n", utf8.RuneCountInString(unicodeString), unicodeString)
	fmt.Printf("Length of string is %d bytes. \n\n", len(unicodeString))

	for i, r := range unicodeString {
		fmt.Printf("Index:%d  Character:%q  Unicode:%d \n", i, r, r)
	}
}
