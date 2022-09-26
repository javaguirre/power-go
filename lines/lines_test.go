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

func TestLinesInputFromArgsValidArgs(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := lines.NewCounter(
		lines.WithInputFromArgs(args),
	)

	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestLinesInputFromArgsEmptyArgs(t *testing.T) {
	t.Parallel()
	args := []string{}
	_, err := lines.NewCounter(
		lines.WithInputFromArgs(args),
	)

	if err.Error() != "not enough arguments" {
		t.Errorf("Should have errored with no args")
	}
}

func TestWordsValidInputOutput(t *testing.T) {
	t.Parallel()
	inputBuf := bytes.NewBufferString("word1 word2 word3\n word4 word5 word6")
	c, err := lines.NewCounter(
		lines.WithInput(inputBuf),
		lines.WithOutput(os.Stdout),
	)
	want := 6

	if err != nil {
		t.Fatal(err)
	}

	got := c.Words()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
