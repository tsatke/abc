package abc

import (
	"fmt"
	"io"
	"sync"
)

const (
	// TimeLayoutSimpleLogger is the time layout that the simple logger
	// uses for its messages.
	TimeLayoutSimpleLogger = "2006-01-02 15:04:05.000"
)

// SimpleLogger is a logger that prints log messages.
// SimpleLoggers are completely safe for concurrent use.
type SimpleLogger struct {
	lvlMux sync.Mutex
	lvl    LogLevel

	clockMux sync.Mutex
	clock    Clock

	outMux sync.Mutex
	out    io.Writer
}

// Print prints the given values with the given log level,
// if and only if the given log level is higher than or
// equal to the one of this logger.
func (s *SimpleLogger) Print(lvl LogLevel, v ...interface{}) {
	if s.IsLevelEnabled(lvl) {
		s.print0(s.prepareMessage(lvl, fmt.Sprint(v...)))
	}
}

// Printf formats and prints the given values with the given log level,
// if and only if the given log level is higher than or
// equal to the one of this logger.
func (s *SimpleLogger) Printf(lvl LogLevel, format string, v ...interface{}) {
	if s.IsLevelEnabled(lvl) {
		s.print0(s.prepareMessage(lvl, fmt.Sprintf(format, v...)))
	}
}

func (s *SimpleLogger) prepareMessage(lvl LogLevel, a string) string {
	return fmt.Sprintf("%v [%-4v] - %v\n", s.clock.Now().Format(TimeLayoutSimpleLogger), lvl.String(), a)
}

func (s *SimpleLogger) print0(a string) {
	io.WriteString(s.out, a)
}

// Verbose prints the given values with log level DEBG,
// if and only if this logger has the verbose log level enabled.
func (s *SimpleLogger) Verbose(v ...interface{}) {
	s.Print(LevelVerbose, v...)
}

// Verbosef formats and prints the given values with log level DEBG,
// if and only if this logger has the verbose log level enabled.
func (s *SimpleLogger) Verbosef(format string, v ...interface{}) {
	s.Printf(LevelVerbose, format, v...)
}

// Debug prints the given values with log level DEBG.
func (s *SimpleLogger) Debug(v ...interface{}) {
	s.Print(LevelDebug, v...)
}

// Debugf formats and prints the given values with log level DEBG.
func (s *SimpleLogger) Debugf(format string, v ...interface{}) {
	s.Printf(LevelDebug, format, v...)
}

// Info prints the given values with log level INFO.
func (s *SimpleLogger) Info(v ...interface{}) {
	s.Print(LevelInfo, v...)
}

// Infof formats and prints the given values with log level INFO.
func (s *SimpleLogger) Infof(format string, v ...interface{}) {
	s.Printf(LevelInfo, format, v...)
}

// Warn prints the given values with log level WARN.
func (s *SimpleLogger) Warn(v ...interface{}) {
	s.Print(LevelWarn, v...)
}

// Warnf formats and prints the given values with log level WARN.
func (s *SimpleLogger) Warnf(format string, v ...interface{}) {
	s.Printf(LevelWarn, format, v...)
}

// Error prints the given values with log level ERR.
func (s *SimpleLogger) Error(v ...interface{}) {
	s.Print(LevelError, v...)
}

// Errorf formats and prints the given values with log level ERR.
func (s *SimpleLogger) Errorf(format string, v ...interface{}) {
	s.Printf(LevelError, format, v...)
}

// Fatal prints the given values with log level FATAL.
// IT DOES NOT TERMINATE THE APPLICATION.
func (s *SimpleLogger) Fatal(v ...interface{}) {
	s.Print(LevelFatal, v...)
}

// Fatalf formats and prints the given values with log level FATAL.
// IT DOES NOT TERMINATE THE APPLICATION.
func (s *SimpleLogger) Fatalf(format string, v ...interface{}) {
	s.Printf(LevelFatal, format, v...)
}

// Level returns the current level of this logger.
func (s *SimpleLogger) Level() LogLevel {
	return s.lvl
}

// SetLevel changes the log level of this logger.
func (s *SimpleLogger) SetLevel(lvl LogLevel) {
	s.lvlMux.Lock()
	defer s.lvlMux.Unlock()
	s.lvl = lvl
}

// IsLevelEnabled returns true if and only if this logger would print
// messages with the given log level.
// False otherwise.
func (s *SimpleLogger) IsLevelEnabled(lvl LogLevel) bool {
	return lvl >= s.lvl
}

// Clock returns the clock of this logger.
func (s *SimpleLogger) Clock() Clock {
	return s.clock
}

// SetClock sets a new clock for this logger.
func (s *SimpleLogger) SetClock(clock Clock) {
	s.clockMux.Lock()
	defer s.clockMux.Unlock()
	s.clock = clock
}

// Out returns the writer of this logger.
func (s *SimpleLogger) Out() io.Writer {
	return s.out
}

// SetOut sets a new writer for this logger.
func (s *SimpleLogger) SetOut(out io.Writer) {
	s.outMux.Lock()
	defer s.outMux.Unlock()
	s.out = out
}
