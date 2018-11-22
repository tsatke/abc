package main

import "gitlab.com/TimSatke/abc"

func main() {
	loggers := []abc.Logger{
		abc.NewSimpleLogger(),
		abc.NewNamedLogger("MyLogger"),
	}

	for _, logger := range loggers {
		logger.SetLevel(abc.LevelDebug)

		logger.Info("I'm alive...")
		logger.Debug("...and can print debug output!")
	}
}
