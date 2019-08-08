package main

import "gitlab.com/TimSatke/abc"

func main() {
	loggers := []abc.Logger{
		abc.NewSimpleLogger(),
		abc.NewNamedLogger("MyLogger"),
		abc.must(abc.NewCustomPatternLogger(`{{.Timestamp}} {{.File}}:{{.Line}} {{.Function}} [{{.Level}}] - {{.Message}}` + "\n")),
	}

	for _, logger := range loggers {
		logger.SetLevel(abc.LevelDebug)

		logger.Info("I'm alive...")
		logger.Debug("...and can print debug output!")
	}
}
