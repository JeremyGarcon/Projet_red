package main

import (
	"os"
	"os/exec"
)

func clear() { // clear le terminal
	var cmd *exec.Cmd
	cmd = exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
