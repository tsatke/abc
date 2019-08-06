package abc

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
	"text/template"
)

const (
	// TimeLayoutCustomPatternLogger is the time layout that the simple logger
	// uses for its messages if no formatting pattern is given.
	TimeLayoutCustomPatternLogger = "2006-01-02 15:04:05.000"
	// CustomPatternLoggerDefaultPattern is the fallback pattern that is used
	// if the pattern compilation or template execution fails.
	CustomPatternLoggerDefaultPattern = "{{.Timestamp}} [{{.Level}}] - {{.Message}}\n"
)

// CustomPatternLogger is a logger that supports a custom
// log pattern.
// The pattern is a go template text.
// If its compilation or execution fails at any point,
// the CustomPatternLoggerDefaultPattern will be used.
// It supports the
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
type CustomPatternLogger struct {
	lvlMux sync.Mutex
	lvl    LogLevel

	clockMux sync.Mutex
	clock    Clock

	outMux sync.Mutex
	out    io.Writer

	pattern  string
	lock     sync.Mutex
	template *template.Template
}

// Print prints the given values with the given log level,
// if and only if the given log level is higher than or
// equal to the one of this logger.
func (l *CustomPatternLogger) Print(lvl LogLevel, v ...interface{}) {
	l.print0(lvl, v...)
}

// Printf formats and prints the given values with the given log level,
// if and only if the given log level is higher than or
// equal to the one of this logger.
func (l *CustomPatternLogger) Printf(lvl LogLevel, format string, v ...interface{}) {
	l.printf0(lvl, format, v...)
}

func (l *CustomPatternLogger) print0(lvl LogLevel, v ...interface{}) {
	if l.IsLevelEnabled(lvl) {
		l.print1(l.prepareMessage(lvl, fmt.Sprint(v...)))
	}
}
func (l *CustomPatternLogger) printf0(lvl LogLevel, format string, v ...interface{}) {
	if l.IsLevelEnabled(lvl) {
		l.print1(l.prepareMessage(lvl, fmt.Sprintf(format, v...)))
	}
}

func (l *CustomPatternLogger) prepareMessage(lvl LogLevel, a string) string {
	if l.template == nil {
		err := l.init()
		if err != nil {
			println(fmt.Sprintf("Failed to initialize logger, using default pattern: %v", err))
			l.pattern = CustomPatternLoggerDefaultPattern
			_ = l.init()
		}
	}

	buf := &bytes.Buffer{}
	err := l.template.Execute(buf, &customPatternLoggerTemplateData{
		clock:   l.clock,
		Level:   fmt.Sprintf("%-4v", lvl.String()),
		Message: a,
	})
	if err != nil {
		println(fmt.Sprintf("Failed to execute template, using default pattern: %v", err))
		l.pattern = CustomPatternLoggerDefaultPattern
		l.template = nil
		return l.prepareMessage(lvl, a) // recursive call, with default pattern, which will be compiled in recursive call
	}
	return buf.String() // TODO(TimSatke) implement message formatting
}

func (l *CustomPatternLogger) init() error {
	// TODO(TimSatke) make safe for concurrent use
	tmpl, err := template.New(l.pattern).Parse(l.pattern)
	if err != nil {
		return err
	}

	l.template = tmpl
	return nil
}

func (l *CustomPatternLogger) print1(a string) {
	io.WriteString(l.out, a)
}

// Verbose prints the given values with log level DEBG,
// if and only if this logger has the verbose log level enabled.
func (l *CustomPatternLogger) Verbose(v ...interface{}) {
	l.print0(LevelVerbose, v...)
}

// Verbosef formats and prints the given values with log level DEBG,
// if and only if this logger has the verbose log level enabled.
func (l *CustomPatternLogger) Verbosef(format string, v ...interface{}) {
	l.printf0(LevelVerbose, format, v...)
}

// Debug prints the given values with log level DEBG.
func (l *CustomPatternLogger) Debug(v ...interface{}) {
	l.print0(LevelDebug, v...)
}

// Debugf formats and prints the given values with log level DEBG.
func (l *CustomPatternLogger) Debugf(format string, v ...interface{}) {
	l.printf0(LevelDebug, format, v...)
}

// Info prints the given values with log level INFO.
func (l *CustomPatternLogger) Info(v ...interface{}) {
	l.print0(LevelInfo, v...)
}

// Infof formats and prints the given values with log level INFO.
func (l *CustomPatternLogger) Infof(format string, v ...interface{}) {
	l.printf0(LevelInfo, format, v...)
}

// Warn prints the given values with log level WARN.
func (l *CustomPatternLogger) Warn(v ...interface{}) {
	l.print0(LevelWarn, v...)
}

// Warnf formats and prints the given values with log level WARN.
func (l *CustomPatternLogger) Warnf(format string, v ...interface{}) {
	l.printf0(LevelWarn, format, v...)
}

// Error prints the given values with log level ERR.
func (l *CustomPatternLogger) Error(v ...interface{}) {
	l.print0(LevelError, v...)
}

// Errorf formats and prints the given values with log level ERR.
func (l *CustomPatternLogger) Errorf(format string, v ...interface{}) {
	l.printf0(LevelError, format, v...)
}

// Fatal prints the given values with log level FATAL.
// IT DOES NOT TERMINATE THE APPLICATION.
func (l *CustomPatternLogger) Fatal(v ...interface{}) {
	l.print0(LevelFatal, v...)
}

// Fatalf formats and prints the given values with log level FATAL.
// IT DOES NOT TERMINATE THE APPLICATION.
func (l *CustomPatternLogger) Fatalf(format string, v ...interface{}) {
	l.printf0(LevelFatal, format, v...)
}

// Level returns the current level of this logger.
func (l *CustomPatternLogger) Level() LogLevel {
	l.lvlMux.Lock()
	defer l.lvlMux.Unlock()

	return l.lvl
}

// SetLevel changes the log level of this logger.
func (l *CustomPatternLogger) SetLevel(lvl LogLevel) {
	l.lvlMux.Lock()
	defer l.lvlMux.Unlock()

	l.lvl = lvl
}

func (s *CustomPatternLogger) SetLevelString(level string) {
	s.SetLevel(ToLogLevel(level))
}

// IsLevelEnabled returns true if and only if this logger would print
// messages with the given log level.
// False otherwise.
func (l *CustomPatternLogger) IsLevelEnabled(lvl LogLevel) bool {
	return lvl >= l.Level()
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

// =======================================================

type customPatternLoggerTemplateData struct {
	clock   Clock
	Level   string
	Message string

	initialized uint32
	callerMux   sync.Mutex
	pc          uintptr
	file        string
	line        int
	function    *runtime.Func
}

func (l *customPatternLoggerTemplateData) Timestamp() string {
	return l.Timestampf(TimeLayoutCustomPatternLogger)
}

func (l *customPatternLoggerTemplateData) Timestampf(layout string) string {
	return l.clock.Now().Format(layout) // the formatted timestamp
}

func (l *customPatternLoggerTemplateData) File() string {
	l.initCallerInfo()
	return l.Filef("short")
}

func (l *customPatternLoggerTemplateData) Filef(mode string) string {
	l.initCallerInfo()
	if mode == "short" {
		name := l.file
		return filepath.Base(name)
	}
	return l.file // calling file
}

func (l *customPatternLoggerTemplateData) Line() int {
	l.initCallerInfo()
	return l.line // calling line number
}

func (l *customPatternLoggerTemplateData) Function() string {
	l.initCallerInfo()
	return l.Functionf("package")
}

func (l *customPatternLoggerTemplateData) Functionf(mode string) string {
	l.initCallerInfo()
	if mode == "short" {
		name := l.function.Name()
		return name[strings.LastIndex(name, ".")+1:]
	} else if mode == "package" {
		return filepath.Base(l.function.Name())
	} else {
		return l.function.Name() // calling package
	}
}

func (l *customPatternLoggerTemplateData) initCallerInfo() {
	if atomic.LoadUint32(&l.initialized) == 1 {
		return
	}

	l.callerMux.Lock()
	defer l.callerMux.Unlock()

	if l.initialized == 0 {
		var ok bool
		l.pc, l.file, l.line, ok = runtime.Caller(18)
		l.function = runtime.FuncForPC(l.pc)

		if ok {
			atomic.AddUint32(&l.initialized, 1)
		}
	}
}
