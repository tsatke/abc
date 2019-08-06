package main

import (
	"fmt"

	"github.com/TimSatke/abc"
)

func main() {
	abc.SetLevel(abc.LevelVerbose) // default level is INFO (abc.LevelInfo)

	abc.Verbose("Hello World")             // prints as DEBG
	abc.Verbosef("fmt: %v", "Hello World") // prints as DEBG

	abc.Debug("Hello World")
	abc.Debugf("fmt: %v", "Hello World")

	abc.Info("Hello World")
	abc.Infof("fmt: %v", "Hello World")

	abc.Warn("Hello World")
	abc.Warnf("fmt: %v", "Hello World")

	abc.Error("Hello World")
	abc.Errorf("fmt: %v", "Hello World")

	abc.Fatal("Hello World")
	abc.Fatalf("fmt: %v", "Hello World")

	fmt.Println("Still running")
}
