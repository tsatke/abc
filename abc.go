package abc

import (
	"os"
	"sync"
)

var (
	mu   sync.Mutex
	root Logger
)

func init() {
	SetRoot(NewSimpleLogger())
}

func Root() Logger {
	return root
}

func SetRoot(lg Logger) {
	mu.Lock()
	defer mu.Unlock()
	root = lg
}

func NewSimpleLogger() Logger {
	return &SimpleLogger{
		lvl:   LevelInfo,
		clock: &realClock{},
		out:   os.Stdout,
	}
}

// "Implementing" abc.Logger

func Print(lvl LogLevel, v ...interface{}) {
	if root.IsLevelEnabled(lvl) {
		root.Print(lvl, v...)
	}
}

func Printf(lvl LogLevel, format string, v ...interface{}) {
	if root.IsLevelEnabled(lvl) {
		root.Printf(lvl, format, v...)
	}
}

func Inspect(v interface{}) {
	panic("Unsupported") // TODO(TimSatke) custom implementation
}

func Verbose(v ...interface{}) {
	Print(LevelVerbose, v...)
}

func Verbosef(format string, v ...interface{}) {
	Printf(LevelVerbose, format, v...)
}

func Debug(v ...interface{}) {
	Print(LevelDebug, v...)
}

func Debugf(format string, v ...interface{}) {
	Printf(LevelDebug, format, v...)
}

func Info(v ...interface{}) {
	Print(LevelInfo, v...)
}

func Infof(format string, v ...interface{}) {
	Printf(LevelInfo, format, v...)
}

func Warn(v ...interface{}) {
	Print(LevelWarn, v...)
}

func Warnf(format string, v ...interface{}) {
	Printf(LevelWarn, format, v...)
}

func Error(v ...interface{}) {
	Print(LevelError, v...)
}

func Errorf(format string, v ...interface{}) {
	Printf(LevelError, format, v...)
}

func Fatal(v ...interface{}) {
	Print(LevelFatal, v...)
}

func Fatalf(format string, v ...interface{}) {
	Printf(LevelFatal, format, v...)
}

func SetLevel(lvl LogLevel) {
	root.SetLevel(lvl)
}

func IsLevelEnabled(lvl LogLevel) bool {
	return root.IsLevelEnabled(lvl)
}
