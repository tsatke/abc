package abc

import (
	"fmt"
	"io"
	"sync"
)

const (
	TimeLayoutSimpleLogger = "2006-01-02 15:04:05.000"
)

type SimpleLogger struct {
	lvlMux sync.Mutex
	lvl    LogLevel

	clockMux sync.Mutex
	clock    Clock

	outMux sync.Mutex
	out    io.Writer
}

func (s *SimpleLogger) Print(lvl LogLevel, v ...interface{}) {
	if s.IsLevelEnabled(lvl) {
		s.print0(s.prepareMessage(lvl, fmt.Sprint(v...)))
	}
}

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

func (s *SimpleLogger) Inspect(v interface{}) {
	panic("Unsupported") // TODO(TimSatke) custom implementation
}

func (s *SimpleLogger) Verbose(v ...interface{}) {
	s.Print(LevelVerbose, v...)
}

func (s *SimpleLogger) Verbosef(format string, v ...interface{}) {
	s.Printf(LevelVerbose, format, v...)
}

func (s *SimpleLogger) Debug(v ...interface{}) {
	s.Print(LevelDebug, v...)
}

func (s *SimpleLogger) Debugf(format string, v ...interface{}) {
	s.Printf(LevelDebug, format, v...)
}

func (s *SimpleLogger) Info(v ...interface{}) {
	s.Print(LevelInfo, v...)
}

func (s *SimpleLogger) Infof(format string, v ...interface{}) {
	s.Printf(LevelInfo, format, v...)
}

func (s *SimpleLogger) Warn(v ...interface{}) {
	s.Print(LevelWarn, v...)
}

func (s *SimpleLogger) Warnf(format string, v ...interface{}) {
	s.Printf(LevelWarn, format, v...)
}

func (s *SimpleLogger) Error(v ...interface{}) {
	s.Print(LevelError, v...)
}

func (s *SimpleLogger) Errorf(format string, v ...interface{}) {
	s.Printf(LevelError, format, v...)
}

func (s *SimpleLogger) Fatal(v ...interface{}) {
	s.Print(LevelFatal, v...)
}

func (s *SimpleLogger) Fatalf(format string, v ...interface{}) {
	s.Printf(LevelFatal, format, v...)
}

func (s *SimpleLogger) Level() LogLevel {
	return s.lvl
}

func (s *SimpleLogger) SetLevel(lvl LogLevel) {
	s.lvlMux.Lock()
	defer s.lvlMux.Unlock()
	s.lvl = lvl
}

func (s *SimpleLogger) IsLevelEnabled(lvl LogLevel) bool {
	return lvl >= s.lvl
}

func (s *SimpleLogger) Clock() Clock {
	return s.clock
}

func (s *SimpleLogger) SetClock(clock Clock) {
	s.clockMux.Lock()
	defer s.clockMux.Unlock()
	s.clock = clock
}

func (s *SimpleLogger) Out() io.Writer {
	return s.out
}

func (s *SimpleLogger) SetOut(out io.Writer) {
	s.outMux.Lock()
	defer s.outMux.Unlock()
	s.out = out
}
