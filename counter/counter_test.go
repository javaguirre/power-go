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
	counter := counter.NewNumCounter()

	counter.Set(5)
	got := counter.Count

	if want != got {
		t.Errorf("NOOOO")
	}
}

// TODO Not sure yet how to make it so it doesn't stop all the execution
// func TestRunCounterwithOneSecondSleep(t *testing.T) {
// 	t.Parallel()
// 	want := ""
// 	counter := counter.NewNumCounter(counter.WithSleepTimeSeconds(1))
// 	fakeTerminal := &bytes.Buffer{}
// 	counter.Output = fakeTerminal

// 	counter.Run()
// 	got := fakeTerminal.String()

// 	if !cmp.Equal(want, got) {
// 		t.Error(cmp.Diff(want, got))
// 	}
// }
