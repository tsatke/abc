package main

import "gitlab.com/TimSatke/abc"

func main() {
	logger := abc.NewSimpleLogger()
	logger.SetLevel(abc.LevelDebug)

	logger.Info("I'm alive!")

	logger.Verbose("This will not print, since Verbose is below Debug")

	logger.Fatal("This will not terminate the application")
	logger.Debug("See? Nothing can stop me!")

	logger.SetLevel(abc.LevelWarn)

	logger.Error("Still printing...")

	abc.SetRoot(logger)
	abc.Warn("Now you can use me globally")
}
