package main

import "gitlab.com/TimSatke/abc"

func main() {
	loggers := []abc.Logger{
		abc.NewSimpleLogger(),
		abc.NewNamedLogger("MyLogger"),
		abc.Must(abc.NewCustomPatternLogger(`{{.Timestamp}} {{.File}}:{{.Line}} {{.Function}} [{{.Level}}] - {{.Message}}` + "\n")),
		abc.NewColoredLogger(abc.NewSimpleLogger()),
	}

	for _, logger := range loggers {
		logger.SetLevel(abc.LevelDebug)

		logger.Info("I'm alive...")
		logger.Debug("...and can print debug output!")
	}
}
