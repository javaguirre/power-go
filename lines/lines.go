package lines

import "os"

// TODO Finish the functional option approach
type option func(counter) counter

func NewCounter(opts ...option) counter {
	c := counter{
		input:  os.Stdin,
		output: os.Stdout,
	}
	for opt, _ := range opts {
		c = opt(c)
	}

	return c
}
