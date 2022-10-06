package counter

import (
	"fmt"
	"io"
	"os"
	"time"
)

type NumCounter struct {
	Count        int
	Output       io.Writer
	SleepSeconds time.Duration
}

func NewNumCounter(opts ...option) *NumCounter {
	counter := &NumCounter{
		Count:        0,
		Output:       os.Stdout,
		SleepSeconds: 0 * time.Second}

	for _, opt := range opts {
		opt(counter)
	}

	return counter
}

type option func(*NumCounter)

func WithSleepTimeSeconds(seconds int) option {
	return func(c *NumCounter) {
		c.SleepSeconds = time.Duration(seconds) * time.Second
	}
}

func (counter *NumCounter) Next() int {
	counter.Count += 1
	return counter.Count
}

func (counter *NumCounter) Set(count int) {
	counter.Count = count
}

// Run all the time, sleep for Sleep duration
func (counter *NumCounter) Run() {
	for {
		counter.Output.Write([]byte(fmt.Sprintf("Counter: %d\n", counter.Next())))
		time.Sleep(counter.SleepSeconds)
	}
}
