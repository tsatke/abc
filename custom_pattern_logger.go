package abc

import (
	"fmt"
	"io"
	"sync"
)

const (
	// TimeLayoutCustomPatternLogger is the time layout that the simple logger
	// uses for its messages if no formatting pattern is given.
	TimeLayoutCustomPatternLogger = "2006-01-02 15:04:05.000"
)

type CustomPatternLogger struct {
	lvlMux sync.Mutex
	lvl    LogLevel

	clockMux sync.Mutex
	clock    Clock

	outMux sync.Mutex
	out    io.Writer

	pattern string
}

// Print prints the given values with the given log level,
// if and only if the given log level is higher than or
// equal to the one of this logger.
func (l *CustomPatternLogger) Print(lvl LogLevel, v ...interface{}) {
	if l.IsLevelEnabled(lvl) {
		l.print0(l.prepareMessage(lvl, fmt.Sprint(v...)))
	}
}

// Printf formats and prints the given values with the given log level,
// if and only if the given log level is higher than or
// equal to the one of this logger.
func (l *CustomPatternLogger) Printf(lvl LogLevel, format string, v ...interface{}) {
	if l.IsLevelEnabled(lvl) {
		l.print0(l.prepareMessage(lvl, fmt.Sprintf(format, v...)))
	}
}

func (l *CustomPatternLogger) prepareMessage(lvl LogLevel, a string) string {
	return "" // TODO(TimSatke) implement message formatting
}

func (l *CustomPatternLogger) print0(a string) {
	io.WriteString(l.out, a)
}

// Inspect prints detailed information about the given value.
// Very verbose, may be slow.
// NOT RECOMMENDED FOR PRODUCTION USE.
func (l *CustomPatternLogger) Inspect(v interface{}) {
	panic("Unsupported") // TODO(TimSatke) custom implementation
}

// Verbose prints the given values with log level DEBG,
// if and only if this logger has the verbose log level enabled.
func (l *CustomPatternLogger) Verbose(v ...interface{}) {
	l.Print(LevelVerbose, v...)
}

// Verbosef formats and prints the given values with log level DEBG,
// if and only if this logger has the verbose log level enabled.
func (l *CustomPatternLogger) Verbosef(format string, v ...interface{}) {
	l.Printf(LevelVerbose, format, v...)
}

// Debug prints the given values with log level DEBG.
func (l *CustomPatternLogger) Debug(v ...interface{}) {
	l.Print(LevelDebug, v...)
}

// Debugf formats and prints the given values with log level DEBG.
func (l *CustomPatternLogger) Debugf(format string, v ...interface{}) {
	l.Printf(LevelDebug, format, v...)
}

// Info prints the given values with log level INFO.
func (l *CustomPatternLogger) Info(v ...interface{}) {
	l.Print(LevelInfo, v...)
}

// Infof formats and prints the given values with log level INFO.
func (l *CustomPatternLogger) Infof(format string, v ...interface{}) {
	l.Printf(LevelInfo, format, v...)
}

// Warn prints the given values with log level WARN.
func (l *CustomPatternLogger) Warn(v ...interface{}) {
	l.Print(LevelWarn, v...)
}

// Warnf formats and prints the given values with log level WARN.
func (l *CustomPatternLogger) Warnf(format string, v ...interface{}) {
	l.Printf(LevelWarn, format, v...)
}

// Error prints the given values with log level ERR.
func (l *CustomPatternLogger) Error(v ...interface{}) {
	l.Print(LevelError, v...)
}

// Errorf formats and prints the given values with log level ERR.
func (l *CustomPatternLogger) Errorf(format string, v ...interface{}) {
	l.Printf(LevelError, format, v...)
}

// Fatal prints the given values with log level FATAL.
// IT DOES NOT TERMINATE THE APPLICATION.
func (l *CustomPatternLogger) Fatal(v ...interface{}) {
	l.Print(LevelFatal, v...)
}

// Fatalf formats and prints the given values with log level FATAL.
// IT DOES NOT TERMINATE THE APPLICATION.
func (l *CustomPatternLogger) Fatalf(format string, v ...interface{}) {
	l.Printf(LevelFatal, format, v...)
}

// Level returns the current level of this logger.
func (l *CustomPatternLogger) Level() LogLevel {
	return l.lvl
}

// SetLevel changes the log level of this logger.
func (l *CustomPatternLogger) SetLevel(lvl LogLevel) {
	l.lvlMux.Lock()
	defer l.lvlMux.Unlock()
	l.lvl = lvl
}

// IsLevelEnabled returns true if and only if this logger would print
// messages with the given log level.
// False otherwise.
func (l *CustomPatternLogger) IsLevelEnabled(lvl LogLevel) bool {
	return lvl >= l.lvl
}

// Clock returns the clock of this logger.
func (l *CustomPatternLogger) Clock() Clock {
	return l.clock
}

// SetClock sets a new clock for this logger.
func (l *CustomPatternLogger) SetClock(clock Clock) {
	l.clockMux.Lock()
	defer l.clockMux.Unlock()
	l.clock = clock
}

// Out returns the writer of this logger.
func (l *CustomPatternLogger) Out() io.Writer {
	return l.out
}

// SetOut sets a new writer for this logger.
func (l *CustomPatternLogger) SetOut(out io.Writer) {
	l.outMux.Lock()
	defer l.outMux.Unlock()
	l.out = out
}
