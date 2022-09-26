package lines_test

import (
	"bytes"
	"testing"

	"github.com/power-go/lines"
)

func TestLines(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("1\n2\n3\n")
	c, err := lines.NewCounter(
		lines.WithInput(inputBuf),
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
