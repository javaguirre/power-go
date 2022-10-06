package main

import (
	"github.com/power-go/counter"
)

func main() {
	counter := counter.NewNumCounter(counter.WithSleepTimeSeconds(1))
	counter.Run()
}
