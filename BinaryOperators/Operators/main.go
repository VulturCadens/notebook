package main

import (
	"encoding/binary"
	"fmt"
)

var (
	bytes []byte = []byte{0xff, 0x00}
	word  uint16 = 0x22aa
)

func setBit(b, m byte) byte {
	b |= m
	return b
}

func clearBit(b, m byte) byte {
	b &^= m
	return b
}

func testBit(b, m byte) bool {
	return (b & m) != 0
}

func main() {
	/*
	 * Decodes a Uint16 from a byte slice.
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
	fmt.Printf(" 1100 XOR \n 0101 -> \n %04b \n\n", XOR)

	/*
	 * Call functions to set, clear and test a bit.
	 */

	var b byte = 0b1000

	fmt.Printf("SetBit    (Set first)    %04b -> %04b \n", b, setBit(b, 0b0001))
	fmt.Printf("ClearBit  (Clear fourth) %04b -> %04b \n", b, clearBit(b, 0b1000))
	fmt.Printf("TestBit   (Fourth set?)  %04b -> %t \n", b, testBit(b, 0b1000))
	fmt.Printf("TestBit   (First set?)   %04b -> %t \n", b, testBit(b, 0b0001))

}
