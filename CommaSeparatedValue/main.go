package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

var filename string = "file.csv"

func main() {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	reader := csv.NewReader(file)

	for {
		row, err := reader.Read()

		if errors.Is(err, io.EOF) {
			break
		}

		for index := range row {
			fmt.Printf("%d %s ", index, row[index])
		}

		fmt.Printf("\n")
	}
}
