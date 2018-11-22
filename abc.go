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

// Root returns the globally used root logger.
func Root() Logger {
	return root
}

// SetRoot changes the globally used root logger.
// This function is safe for concurrent use.
func SetRoot(lg Logger) {
	mu.Lock()
	defer mu.Unlock()
	root = lg
}

// NewSimpleLogger returns a new abc.SimpleLogger,
// which is ready to use.
// The default log level is INFO and can be changed with
//
//	logger.SetLevel(abc.LevelInfo)
//
// The logger prints to os.Stdout by default.
// The output writer can be changed with
//
//	logger.SetOut(os.Stdout)
func NewSimpleLogger() WriterLogger {
	return &SimpleLogger{
		lvl:   LevelInfo,
		clock: &realClock{},
		out:   os.Stdout,
	}
}

func NewNamedLogger(name string) WriterLogger {
	return &NamedLogger{
		lvl:   LevelInfo,
		clock: &realClock{},
		out:   os.Stdout,
		name:  name,
	}
}

// "Implementing" abc.Logger

// Print delegates to the root logger, if and only if
// the given level is enabled by the root logger.
func Print(lvl LogLevel, v ...interface{}) {
	if root.IsLevelEnabled(lvl) {
		root.Print(lvl, v...)
	}
}

// Print delegates to the root logger, if and only if
// the given level is enabled by the root logger.
func Printf(lvl LogLevel, format string, v ...interface{}) {
	if root.IsLevelEnabled(lvl) {
		root.Printf(lvl, format, v...)
	}
}

// Inspect is coming soon...
func Inspect(v interface{}) {
	panic("Unsupported") // TODO(TimSatke) custom implementation
}

// Verbose prints the given values with log level DEBG,
// but the root logger must have verbose log levels enabled
// to show any output.
func Verbose(v ...interface{}) {
	Print(LevelVerbose, v...)
}

// Verbosef formats and prints the given values with log level DEBG,
// but the root logger must have verbose log levels enabled
// to show any output.
func Verbosef(format string, v ...interface{}) {
	Printf(LevelVerbose, format, v...)
}

// Debug prints the given values with log level DEBG.
func Debug(v ...interface{}) {
	Print(LevelDebug, v...)
}

// Debugf formats and prints the given values with log level DEBG.
func Debugf(format string, v ...interface{}) {
	Printf(LevelDebug, format, v...)
}

// Info prints the given values with log level INFO.
func Info(v ...interface{}) {
	Print(LevelInfo, v...)
}

// Infof formats and prints the given values with log level INFO.
func Infof(format string, v ...interface{}) {
	Printf(LevelInfo, format, v...)
}

// Warn prints the given values with log level WARN.
func Warn(v ...interface{}) {
	Print(LevelWarn, v...)
}

// Warnf formats and prints the given values with log level WARN.
func Warnf(format string, v ...interface{}) {
	Printf(LevelWarn, format, v...)
}

// Error prints the given values with log level ERR.
func Error(v ...interface{}) {
	Print(LevelError, v...)
}

// Errorf formats and prints the given values with log level ERR.
func Errorf(format string, v ...interface{}) {
	Printf(LevelError, format, v...)
}

// Fatal prints the given values with log level FATAL.
func Fatal(v ...interface{}) {
	Print(LevelFatal, v...)
}

// Fatalf formats and prints the given values with log level FATAL.
func Fatalf(format string, v ...interface{}) {
	Printf(LevelFatal, format, v...)
}

// SetLevel sets a new log level for the root logger.
func SetLevel(lvl LogLevel) {
	root.SetLevel(lvl)
}

// IsLevelEnabled returns true if and only if the root logger would print
// messages with the given log level.
// False otherwise.
func IsLevelEnabled(lvl LogLevel) bool {
	return root.IsLevelEnabled(lvl)
}
