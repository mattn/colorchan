package main

import (
	"github.com/mattn/go-colorable"
	"io"
	"os"
)

func main() {
	io.Copy(colorable.NewColorableStdout(), os.Stdin)
}
