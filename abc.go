package abc

import (
	"fmt"
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
		lvl: LevelInfo,
		clk: &realClock{},
		out: os.Stdout,
	}
}

// NewNamedLogger returns a new abc.NamedLogger,
// which is ready to use.
// The default log level is INFO and can be changed with
//
//	logger.SetLevel(abc.LevelInfo)
//
// The logger prints to os.Stdout by default.
// The output writer can be changed with
//
//	logger.SetOut(os.Stdout)
func NewNamedLogger(name string) WriterLogger {
	return &NamedLogger{
		lvl:  LevelInfo,
		clk:  &realClock{},
		out:  os.Stdout,
		name: name,
	}
}

// NewCustomPatternLogger returns a new abc.CustomPatternLogger,
// which was initialized and thus is ready to use.
// If an error occurs during the initialization, that error is
// returned.
//
// The given pattern is a go template text and supports the
// following operations:
//
//	{{.Level}} // the level of the message
// Level prints the level of the log message, at least 4 characters.
//
//	{{.Message}} // the message to be printed
// Message prints the log message that should be printed.
// If the message or the pattern doesn't end with a line break,
// no line break will be printed.
//
//	{{.Timestamp}} or {{.Timestampf "2006-01-02 03:04:05PM"}} // time.Time.Format's layout is used
// Timestampf takes a string argument, which will be used for formatting
// the timestamp in the log message. The reference time is the same as in
// time.Time's function "Format".
// Timestamp uses the layout "2006-01-02 15:04:05.000".
//
//	{{.File}} or {{.Filef "short"}} // one of "short", "full" (anything different will be interpreted as "full")
// Filef "short" prints only the filename, while Filef "full"
// prints the file's absolute path.
//
//	{{.Line}} // prints the line of the output call
// Line only prints the line number.
//
//	{{.Function}} or {{.Functionf "package"}} // one of "short", "package", "full" (anything different will be interpreted as "full")
// Functionf "short" prints only the function name, while Functionf "package"
// will print <package>.<function>.
// Functionf "full" will print <full_package>.<function>, e.g. "github.com/TimSatke/abc.main".
//
// Example:
//
//	{{.Timestamp}} {{.Filef "short"}}:{{.Line}} {{.Functionf "package"}} [{{.Level}}] - {{.Message}}\n
//
// will print something like
//
//	2018-11-24 15:26:44.453 main.go:16 main.main [INFO] - Hello World!
//	<line break>
func NewCustomPatternLogger(pattern string) (WriterLogger, error) {
	logger := &CustomPatternLogger{
		lvl:     LevelInfo,
		clk:     &realClock{},
		out:     os.Stdout,
		pattern: pattern,
	}
	err := logger.init()
	return logger, err
}

// NewColoredLogger creates a wrapper for a given WriterLogger.
// Depending on the level that should be printed, this wrapper
// will prepend an ANSI-color code to the wrapped loggers
// output writer and will then call the respective output method.
// Please notice that this function does not add a decorator to
// the given logger, but creates a wrapper, which must be used
// for colors to show up.
func NewColoredLogger(wrapped WriterLogger) WriterLogger {
	return &ColoredLogger{
		wrapped: wrapped,
	}
}

// must panics, if the given error is not nil.
// It returns the unmodified given logger otherwise.
func must(logger Logger, err error) Logger {
	if err != nil {
		panic(fmt.Errorf("must: %v", err))
	}

	return logger
}

// "Implementing" abc.Logger

// Print delegates to the root logger, if and only if
// the given level is enabled by the root logger.
func Print(lvl LogLevel, v ...interface{}) {
	if root.IsLevelEnabled(lvl) {
		root.Print(lvl, v...)
	}
}

// Printf delegates to the root logger, if and only if
// the given level is enabled by the root logger.
func Printf(lvl LogLevel, format string, v ...interface{}) {
	if root.IsLevelEnabled(lvl) {
		root.Printf(lvl, format, v...)
	}
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
