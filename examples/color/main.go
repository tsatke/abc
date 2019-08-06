package main

import "github.com/TimSatke/abc"

func main() {
	l := abc.NewSimpleLogger()
	logger := abc.NewColoredLogger(l)
	logger.SetLevel(abc.LevelVerbose)
	logger.Verbose("Hello World!")
	logger.Verbosef("fmt: Hello %v!", "World")
	logger.Debug("Hello World!")
	logger.Debugf("fmt: Hello %v!", "World")
	logger.Info("Hello World!")
	logger.Infof("fmt: Hello %v!", "World")
	logger.Warn("Hello World!")
	logger.Warnf("fmt: Hello %v!", "World")
	logger.Error("Hello World!")
	logger.Errorf("fmt: Hello %v!", "World")
	logger.Fatal("Hello World!")
	logger.Fatalf("fmt: Hello %v!", "World")
}
