package counter_test

import (
	"testing"

	"github.com/power-go/counter"
)

func TestSetCounterToTwo(t *testing.T) {
	t.Parallel()
	want := 2
	counter := counter.NewNumCounter()

	counter.Next()
	got := counter.Next()

	if want != got {
		t.Errorf("NOOOO")
	}
}

func TestSetCounterToAnyNumber(t *testing.T) {
	t.Parallel()
	want := 5
	counter := counter.NewCounter()

	counter.Set(5)
	got := counter.Count

	if want != got {
		t.Errorf("NOOOO")
	}
}
