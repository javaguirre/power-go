package counter

import "time"

type NumCounter struct {
	Count int
	Sleep time.Duration
}

func NewNumCounter() *NumCounter {
	return &NumCounter{Count: 0, Sleep: 0 * time.Second}
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
}
