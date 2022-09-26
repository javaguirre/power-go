package main

import (
	"fmt"

	"github.com/power-go/lines"
)

func main() {
	wordCount, err := lines.Words()

	if err != nil {
		panic("There was an error in the program")
	}

	fmt.Println(wordCount)
}
