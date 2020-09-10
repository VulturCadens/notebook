package main

import (
	"encoding/binary"
	"fmt"
)

var (
	bytes []byte = []byte{0xff, 0x00}
	word  uint16 = 0x22aa
)

func main() {
	/*
	 * Decodes an Uint16 from a byte slice.
	 */

	wordLittleEndian := binary.LittleEndian.Uint16(bytes[0:])
	wordBigEndian := binary.BigEndian.Uint16(bytes[0:])

	fmt.Printf("Little-endian: %04x \n", wordLittleEndian)
	fmt.Printf("Big-endian: %04x \n\n", wordBigEndian)

	/*
	 * Encodes a Uint16 into a byte slice.
	 */

	twoBytes := make([]byte, 2)

	binary.LittleEndian.PutUint16(twoBytes[0:], word)
	fmt.Printf("Little-endian: % x \n", twoBytes)

	binary.BigEndian.PutUint16(twoBytes[0:], word)
	fmt.Printf("Big-endian: % x \n\n", twoBytes)

	/*
	 * Bitwise operators
	 */

	AND := 0b1100 & 0b0101
	OR := 0b1100 | 0b0101
	XOR := 0b1100 ^ 0b0101

	fmt.Printf(" 1100 AND \n 0101 -> \n %04b \n\n", AND)
	fmt.Printf(" 1100 OR \n 0101 -> \n %04b \n\n", OR)
	fmt.Printf(" 1100 XOR \n 0101 -> \n %04b \n", XOR)
}
