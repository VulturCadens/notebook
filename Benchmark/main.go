package main

import (
	"benchmark/benchJSON"
	"benchmark/benchPB"
	"flag"
	"fmt"
	"io/ioutil"
)

var write bool

func main() {
	flag.BoolVar(&write, "write", false, "Write BOOK.json and BOOK.protobuf files.")
	flag.Parse()

	protobufBook := benchPB.StructToProtobuf()
	jsonBook := benchJSON.StructToJson()

	if write {
		if err := ioutil.WriteFile("BOOK.json", jsonBook, 0644); err != nil {
			panic(err)
		}
		if err := ioutil.WriteFile("BOOK.protobuf", protobufBook, 0644); err != nil {
			panic(err)
		}
	} else {
		fmt.Printf("Protocol Buffers: %d bytes\n", len(protobufBook))
		fmt.Printf("JSON: %d bytes\n", len(jsonBook))
	}
}
