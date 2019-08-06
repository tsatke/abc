package main

import (
	"io"
	"os"

	"github.com/TimSatke/abc"
)

func main() {
	file, _ := os.Create("my.file")

	logger := abc.NewSimpleLogger()
	logger.SetOut(io.MultiWriter(file, os.Stdout)) // writes to file and stdout
	logger.Info("Some piece of information")
}
