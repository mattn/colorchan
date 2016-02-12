package main

import (
	"github.com/mattn/go-colorable"
	"io"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) == 1 {
		stdout := colorable.NewColorableStdout()
		io.Copy(stdout, os.Stdin)
		stdout.Write([]byte("\x1b[39m"))
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
			cmd.Stdout.Write([]byte("\x1b[39m"))
			cmd.Stderr.Write([]byte("\x1b[39m"))
			os.Exit(1)
		}
	}
}
