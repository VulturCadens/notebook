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

	defer file.Close()

	reader := csv.NewReader(file)

	for {
		row, err := reader.Read()

		if errors.Is(err, io.EOF) {
			break
		}

		/*
			for index := range row {
				fmt.Printf("%d %s", index, row[index])
			}

			fmt.Printf("\n")
		*/

		fmt.Printf("%-25s %-20s %4s \n", row[0], row[1], row[2])
	}
}
