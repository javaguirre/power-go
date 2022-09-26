package hello_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/power-go/hello"
)

func TestPrintsHelloMessageToTerminal(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	want := "Hello, Pepe\n"
	p := &hello.Printer{
		Output: fakeTerminal,
	}

	p.Print("Pepe")
	got := fakeTerminal.String()

	if want != got {
		t.Errorf("NOOOO")
	}
}

func TestPrintFixedRightMinutesAndHour(t *testing.T) {
	t.Parallel()
	fakeTerminal := &bytes.Buffer{}
	want := "It's 10 minutes past 8"
	myTime, _ := time.Parse("2006 Jan 02 15:04:05", "2012 Dec 07 08:10:30.918273645")

	hello.PrintTime(fakeTerminal, &myTime)
	got := fakeTerminal.String()

	if want != got {
		t.Errorf("NOO")
	}
}
