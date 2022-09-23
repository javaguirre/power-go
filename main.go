package main

import (
	"os"

	"github.com/power-go/hello"
)

func main() {
	hello.PrintTo(os.Stdout, "Hey")
}
