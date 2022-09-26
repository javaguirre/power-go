package hello

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Printer struct {
	Output io.Writer
}

func NewPrinter() *Printer {
	return &Printer{
		Output: os.Stdout,
	}
}

func (p *Printer) Print(name string) {
	message := "Hello, " + name + "\n"
	p.Output.Write([]byte(message))
}

func PrintTime(writer io.Writer, theTime *time.Time) {
	message := fmt.Sprintf("It's %d minutes past %d", theTime.Minute(), theTime.Hour())
	writer.Write([]byte(message))
}
