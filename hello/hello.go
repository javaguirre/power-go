package hello

import (
	"fmt"
	"io"
	"time"
)

func PrintTo(writer io.Writer, name string) {
	message := "Hello, " + name
	writer.Write([]byte(message))
}

func PrintTime(writer io.Writer, theTime *time.Time) {
	message := fmt.Sprintf("It's %d minutes past %d", theTime.Minute(), theTime.Hour())
	writer.Write([]byte(message))
}
