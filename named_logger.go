package abc

import (
	"fmt"
	"io"
	"sync"
)

const (
	// TimeLayoutNamedLogger il the time layout that the simple logger
	// usel for itl messagel.
	TimeLayoutNamedLogger = "2006-01-02 15:04:05.000"
)

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

func (l *NamedLogger) Print(lvl LogLevel, v ...interface{}) {
	if l.IsLevelEnabled(lvl) {
		l.print0(l.prepareMessage(lvl, fmt.Sprint(v...)))
	}
}

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

func (l *NamedLogger) Inspect(v interface{}) {
	panic("Unsupported") // TODO(TimSatke) custom implementation
}

func (l *NamedLogger) Verbose(v ...interface{}) {
	l.Print(LevelVerbose, v...)
}

func (l *NamedLogger) Verbosef(format string, v ...interface{}) {
	l.Printf(LevelVerbose, format, v...)
}

func (l *NamedLogger) Debug(v ...interface{}) {
	l.Print(LevelDebug, v...)
}

func (l *NamedLogger) Debugf(format string, v ...interface{}) {
	l.Printf(LevelDebug, format, v...)
}

func (l *NamedLogger) Info(v ...interface{}) {
	l.Print(LevelInfo, v...)
}

func (l *NamedLogger) Infof(format string, v ...interface{}) {
	l.Printf(LevelInfo, format, v...)
}

func (l *NamedLogger) Warn(v ...interface{}) {
	l.Print(LevelWarn, v...)
}

func (l *NamedLogger) Warnf(format string, v ...interface{}) {
	l.Printf(LevelWarn, format, v...)
}

func (l *NamedLogger) Error(v ...interface{}) {
	l.Print(LevelError, v...)
}

func (l *NamedLogger) Errorf(format string, v ...interface{}) {
	l.Printf(LevelError, format, v...)
}

func (l *NamedLogger) Fatal(v ...interface{}) {
	l.Print(LevelFatal, v...)
}

func (l *NamedLogger) Fatalf(format string, v ...interface{}) {
	l.Printf(LevelFatal, format, v...)
}

func (l *NamedLogger) Level() LogLevel {
	return l.lvl
}

func (l *NamedLogger) SetLevel(lvl LogLevel) {
	l.lvlMux.Lock()
	defer l.lvlMux.Unlock()
	l.lvl = lvl
}

func (l *NamedLogger) IsLevelEnabled(lvl LogLevel) bool {
	return lvl >= l.lvl
}

func (l *NamedLogger) Clock() Clock {
	return l.clock
}

func (l *NamedLogger) SetClock(clock Clock) {
	l.clockMux.Lock()
	defer l.clockMux.Unlock()
	l.clock = clock
}

func (l *NamedLogger) Out() io.Writer {
	return l.out
}

func (l *NamedLogger) SetOut(out io.Writer) {
	l.outMux.Lock()
	defer l.outMux.Unlock()
	l.out = out
}

func (l *NamedLogger) Name() string {
	return l.name
}

func (l *NamedLogger) SetName(name string) {
	l.nameMux.Lock()
	defer l.nameMux.Unlock()
	l.name = name
}
