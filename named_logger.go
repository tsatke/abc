package abc

import (
	"fmt"
	"io"
	"sync"
)

const (
	// TimeLayoutNamedLogger is the time layout that the named logger
	// uses for its messages.
	TimeLayoutNamedLogger = "2006-01-02 15:04:05.000"
)

// NamedLogger is a logger that has and prints a name
// in its log messages.
// NamedLoggers are completely safe for concurrent use.
type NamedLogger struct {
	lvlMux sync.Mutex
	lvl    LogLevel

	clockMux sync.Mutex
	clock    Clock

	outMux sync.Mutex
	out    io.Writer

	nameMux sync.Mutex
	name    string
}

// Print prints the given values with the given log level,
// if and only if the given log level is higher than or
// equal to the one of this logger.
func (l *NamedLogger) Print(lvl LogLevel, v ...interface{}) {
	if l.IsLevelEnabled(lvl) {
		l.print0(l.prepareMessage(lvl, fmt.Sprint(v...)))
	}
}

// Printf formats and prints the given values with the given log level,
// if and only if the given log level is higher than or
// equal to the one of this logger.
func (l *NamedLogger) Printf(lvl LogLevel, format string, v ...interface{}) {
	if l.IsLevelEnabled(lvl) {
		l.print0(l.prepareMessage(lvl, fmt.Sprintf(format, v...)))
	}
}

func (l *NamedLogger) prepareMessage(lvl LogLevel, a string) string {
	return fmt.Sprintf("%v <%-v> [%-4v] - %v\n", l.clock.Now().Format(TimeLayoutNamedLogger), l.name, lvl.String(), a)
}

func (l *NamedLogger) print0(a string) {
	io.WriteString(l.out, a)
}

// Verbose prints the given values with log level DEBG,
// if and only if this logger has the verbose log level enabled.
func (l *NamedLogger) Verbose(v ...interface{}) {
	l.Print(LevelVerbose, v...)
}

// Verbosef formats and prints the given values with log level DEBG,
// if and only if this logger has the verbose log level enabled.
func (l *NamedLogger) Verbosef(format string, v ...interface{}) {
	l.Printf(LevelVerbose, format, v...)
}

// Debug prints the given values with log level DEBG.
func (l *NamedLogger) Debug(v ...interface{}) {
	l.Print(LevelDebug, v...)
}

// Debugf formats and prints the given values with log level DEBG.
func (l *NamedLogger) Debugf(format string, v ...interface{}) {
	l.Printf(LevelDebug, format, v...)
}

// Info prints the given values with log level INFO.
func (l *NamedLogger) Info(v ...interface{}) {
	l.Print(LevelInfo, v...)
}

// Infof formats and prints the given values with log level INFO.
func (l *NamedLogger) Infof(format string, v ...interface{}) {
	l.Printf(LevelInfo, format, v...)
}

// Warn prints the given values with log level WARN.
func (l *NamedLogger) Warn(v ...interface{}) {
	l.Print(LevelWarn, v...)
}

// Warnf formats and prints the given values with log level WARN.
func (l *NamedLogger) Warnf(format string, v ...interface{}) {
	l.Printf(LevelWarn, format, v...)
}

// Error prints the given values with log level ERR.
func (l *NamedLogger) Error(v ...interface{}) {
	l.Print(LevelError, v...)
}

// Errorf formats and prints the given values with log level ERR.
func (l *NamedLogger) Errorf(format string, v ...interface{}) {
	l.Printf(LevelError, format, v...)
}

// Fatal prints the given values with log level FATAL.
// IT DOES NOT TERMINATE THE APPLICATION.
func (l *NamedLogger) Fatal(v ...interface{}) {
	l.Print(LevelFatal, v...)
}

// Fatalf formats and prints the given values with log level FATAL.
// IT DOES NOT TERMINATE THE APPLICATION.
func (l *NamedLogger) Fatalf(format string, v ...interface{}) {
	l.Printf(LevelFatal, format, v...)
}

// Level returns the current level of this logger.
func (l *NamedLogger) Level() LogLevel {
	return l.lvl
}

// SetLevel changes the log level of this logger.
func (l *NamedLogger) SetLevel(lvl LogLevel) {
	l.lvlMux.Lock()
	defer l.lvlMux.Unlock()
	l.lvl = lvl
}

// IsLevelEnabled returns true if and only if this logger would print
// messages with the given log level.
// False otherwise.
func (l *NamedLogger) IsLevelEnabled(lvl LogLevel) bool {
	return lvl >= l.lvl
}

// Clock returns the clock of this logger.
func (l *NamedLogger) Clock() Clock {
	return l.clock
}

// SetClock sets a new clock for this logger.
func (l *NamedLogger) SetClock(clock Clock) {
	l.clockMux.Lock()
	defer l.clockMux.Unlock()
	l.clock = clock
}

// Out returns the writer of this logger.
func (l *NamedLogger) Out() io.Writer {
	return l.out
}

// SetOut sets a new writer for this logger.
func (l *NamedLogger) SetOut(out io.Writer) {
	l.outMux.Lock()
	defer l.outMux.Unlock()
	l.out = out
}

// Name returns the name of this logger.
func (l *NamedLogger) Name() string {
	return l.name
}

// SetName sets a new name for this logger.
func (l *NamedLogger) SetName(name string) {
	l.nameMux.Lock()
	defer l.nameMux.Unlock()
	l.name = name
}
