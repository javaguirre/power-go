package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("/usr/bin/pmset", "-g", "ps")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
