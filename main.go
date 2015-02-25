package main

import (
	"github.com/mattn/go-colorable"
	"io"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) == 0 {
		io.Copy(colorable.NewColorableStdout(), os.Stdin)
	} else {
		cmd := exec.Command(os.Args[1], os.Args[2:]...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = colorable.NewColorableStdout()
		cmd.Stderr = colorable.NewColorableStderr()
		err := cmd.Run()
		if err != nil {
			if cmd.ProcessState != nil {
				os.Exit(cmd.ProcessState.Sys().(syscall.WaitStatus).ExitStatus())
			}
			os.Exit(1)
		}
	}
}
