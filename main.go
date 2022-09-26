package main

import (
	"fmt"

	"github.com/power-go/lines"
)

func main() {
	lineCount, err := lines.Lines()
	if err != nil {
		panic("There was an error in the program")
	}

	fmt.Println(lineCount)
}
