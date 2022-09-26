package lines_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/power-go/lines"
)

func TestLinesValidInputOutput(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3\n")
	c, err := lines.NewCounter(
		lines.WithInput(inputBuf),
		lines.WithOutput(os.Stdout),
	)
	want := 3

	if err != nil {
		t.Fatal(err)
	}

	got := c.Lines()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestLinesInvalidInputOutput(t *testing.T) {
	t.Parallel()
	_, err := lines.NewCounter(
		lines.WithInput(nil),
		lines.WithOutput(nil),
	)

	if err == nil {
		t.Errorf("Wanted error")
	}
}

func TestLinesValidInputInvalidOutput(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3\n")
	_, err := lines.NewCounter(
		lines.WithInput(inputBuf),
		lines.WithOutput(nil),
	)

	if err == nil {
		t.Errorf("Wanted error")
	}
}
