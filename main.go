package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// hello.NewPrinter().Print("Pepe")
	// fmt.Println(strconv.Itoa(counter.NewCounter().Next()))

	lines := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines++
	}

	fmt.Println(lines)
}
