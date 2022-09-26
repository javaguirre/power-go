package main

import (
	"fmt"
)

func main() {
	wordCount, err := lines.Words()

	if err != nil {
		panic("There was an error in the program")
	}

	fmt.Println(wordCount)
}
