package abc

import (
	"bytes"
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
	if _, ok := Root().(*SimpleLogger); !ok {
		t.Error("Expected root logger to be of type *SimpleLogger, but was not.")
	}
}

func TestSetRoot(t *testing.T) {
	assert := assert.New(t)

	temp := Root()      // save original root logger
	defer SetRoot(temp) // cleanup

	// no output during tests
	l := &SimpleLogger{
		clk: &mockClock{},
		lvl: LevelInfo,
		out: ioutil.Discard,
	}
	SetRoot(l)

	Info("foo") // should write into ioutil.Discard

	buf := &bytes.Buffer{}

	logger := &SimpleLogger{
		clk: &mockClock{},
		lvl: LevelInfo,
		out: buf,
	}
	SetRoot(logger) // set new root logger

	Info("abc")

	expected := "0001-01-01 00:00:00.000 [INFO] - abc\n"
	assert.Equal(expected, buf.String(), "Wrong output")
}

func TestNewSimpleLogger(t *testing.T) {
	assert := assert.New(t)

	l := NewSimpleLogger()

	// check type
	logger, ok := l.(*SimpleLogger)
	assert.True(ok, "Logger was expected to be of type *SimpleLogger, but was not.")

	// check default level
	assert.Equal(logger.lvl, LevelInfo, "Expected level to be INFO")
	// check default out
	assert.Equal(logger.out, os.Stdout)
	// check default clock type (must be real clock)
	_, ok = logger.clk.(*realClock)
	assert.True(ok, "Clock was expected to be of type *realClock, but was not.")
}

func TestNewNamedLogger(t *testing.T) {
	assert := assert.New(t)

	name := "MyLogger"

	l := NewNamedLogger(name)

	// check type
	logger, ok := l.(*NamedLogger)
	assert.True(ok, "Logger was expected to be of type *NamedLogger, but was not.")

	// check default level
	assert.Equal(logger.lvl, LevelInfo, "Expected level to be INFO")
	// check default out
	assert.Equal(logger.out, os.Stdout)
	// check default clock type (must be real clock)
	_, ok = logger.clk.(*realClock)
	assert.True(ok, "Clock was expected to be of type *realClock, but was not.")
	// check name
	assert.Equalf(logger.name, name, "Expected name of logger to be '%v', but was '%v'.", name, logger.name)
}

func TestNewCustomPatternLogger(t *testing.T) {
	assert := assert.New(t)

	pattern := "message: {{.Message}}"

	l, err := NewCustomPatternLogger(pattern)
	assert.NoError(err, "Unable to create a new CustomPatternLogger.")

	// check type
	logger, ok := l.(*CustomPatternLogger)
	assert.True(ok, "Logger was expected to be of type *CustomPatternLogger, but was not.")

	// check default level
	assert.Equal(logger.lvl, LevelInfo, "Expected level to be INFO")
	// check default out
	assert.Equal(logger.out, os.Stdout)
	// check default clock type (must be real clock)
	_, ok = logger.clk.(*realClock)
	assert.True(ok, "Clock was expected to be of type *realClock, but was not.")
	// check pattern
	assert.Equalf(logger.pattern, pattern, "Expected pattern of logger to be '%v', but was '%v'.", pattern, logger.pattern)
}

func TestNewColoredLogger(t *testing.T) {
	assert := assert.New(t)

	tmp := NewSimpleLogger()
	l := NewColoredLogger(tmp)

	// check type
	logger, ok := l.(*ColoredLogger)
	assert.True(ok, "Logger was expected to be of type *ColoredLogger, but was not.")

	// check wrapped logger
	assert.Equal(logger.wrapped, tmp, "Wrapped logger is not the logger that was given.")
}

func Test_All_Outputs(t *testing.T) {
	assert := assert.New(t)

	expectations := []string{
		"0001-01-01 00:00:00.000 [DEBG] - verbose: abc\n",
		"0001-01-01 00:00:00.000 [DEBG] - verbose: fmt: abc\n",
		"0001-01-01 00:00:00.000 [DEBG] - abc\n",
		"0001-01-01 00:00:00.000 [DEBG] - fmt: abc\n",
		"0001-01-01 00:00:00.000 [INFO] - abc\n",
		"0001-01-01 00:00:00.000 [INFO] - fmt: abc\n",
		"0001-01-01 00:00:00.000 [WARN] - abc\n",
		"0001-01-01 00:00:00.000 [WARN] - fmt: abc\n",
		"0001-01-01 00:00:00.000 [ERR ] - abc\n",
		"0001-01-01 00:00:00.000 [ERR ] - fmt: abc\n",
		"0001-01-01 00:00:00.000 [FATAL] - abc\n",
		"0001-01-01 00:00:00.000 [FATAL] - fmt: abc\n",
	}
	cnt := 0

	buf := &bytes.Buffer{}

	logger := &SimpleLogger{
		clk: &mockClock{},
		lvl: LevelVerbose,
		out: buf,
	}
	SetRoot(logger)

	check := func() {
		defer func() {
			cnt++
		}()
		defer buf.Reset()

		if cnt >= len(expectations) {
			panic("No more expectations")
		}

		assert.Equal(expectations[cnt], buf.String(), "Wrong output")
	}

	// actual test flow

	Verbose("verbose: abc")
	check()
	Verbosef("verbose: fmt: %v", "abc")
	check()
	Debug("abc")
	check()
	Debugf("fmt: %v", "abc")
	check()
	Info("abc")
	check()
	Infof("fmt: %v", "abc")
	check()
	Warn("abc")
	check()
	Warnf("fmt: %v", "abc")
	check()
	Error("abc")
	check()
	Errorf("fmt: %v", "abc")
	check()
	Fatal("abc")
	check()
	Fatalf("fmt: %v", "abc")
	check()
}

func TestMustNoPanic(t *testing.T) {
	l := NewSimpleLogger()
	foo := Must(l, nil)
	assert.Equal(l, foo, "Must must return the passed logger")
}

func TestMustWithPanic(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Error("No panic")
		}
	}()

	_ = Must(nil, errors.New("This error was panicked intentionally"))
}

func TestSetLevel(t *testing.T) {
	assert := assert.New(t)

	temp := Root()      // save original root logger
	defer SetRoot(temp) // cleanup

	buf := &bytes.Buffer{}

	logger := &SimpleLogger{
		clk: &mockClock{},
		lvl: LevelInfo,
		out: buf,
	}
	SetRoot(logger) // set new root logger

	Info("abc")

	expected := "0001-01-01 00:00:00.000 [INFO] - abc\n"
	assert.Equal(expected, buf.String(), "Wrong output")

	buf.Reset() // reset buffer

	SetLevel(LevelWarn) // set new log level

	Info("foo")

	assert.Empty(buf.Bytes(), "Buffer received input although log level should be higher than the one printed.")
}

func TestIsLevelEnabled(t *testing.T) {
	assert := assert.New(t)

	defer SetLevel(LevelInfo) // cleanup

	SetLevel(LevelVerbose)
	assert.True(IsLevelEnabled(LevelVerbose), "Verbose level must be enabled")
	assert.True(IsLevelEnabled(LevelDebug), "Debug level must be enabled")
	assert.True(IsLevelEnabled(LevelInfo), "Info level must be enabled")
	assert.True(IsLevelEnabled(LevelWarn), "Warn level must be enabled")
	assert.True(IsLevelEnabled(LevelError), "Error level must be enabled")
	assert.True(IsLevelEnabled(LevelFatal), "Fatal level must be enabled")

	SetLevel(LevelDebug)
	assert.False(IsLevelEnabled(LevelVerbose), "Verbose level must not be enabled")
	assert.True(IsLevelEnabled(LevelDebug), "Debug level must be enabled")
	assert.True(IsLevelEnabled(LevelInfo), "Info level must be enabled")
	assert.True(IsLevelEnabled(LevelWarn), "Warn level must be enabled")
	assert.True(IsLevelEnabled(LevelError), "Error level must be enabled")
	assert.True(IsLevelEnabled(LevelFatal), "Fatal level must be enabled")

	SetLevel(LevelInfo)
	assert.False(IsLevelEnabled(LevelVerbose), "Verbose level must not be enabled")
	assert.False(IsLevelEnabled(LevelDebug), "Debug level must not be enabled")
	assert.True(IsLevelEnabled(LevelInfo), "Info level must be enabled")
	assert.True(IsLevelEnabled(LevelWarn), "Warn level must be enabled")
	assert.True(IsLevelEnabled(LevelError), "Error level must be enabled")
	assert.True(IsLevelEnabled(LevelFatal), "Fatal level must be enabled")

	SetLevel(LevelWarn)
	assert.False(IsLevelEnabled(LevelVerbose), "Verbose level must not be enabled")
	assert.False(IsLevelEnabled(LevelDebug), "Debug level must not be enabled")
	assert.False(IsLevelEnabled(LevelInfo), "Info level must not be enabled")
	assert.True(IsLevelEnabled(LevelWarn), "Warn level must be enabled")
	assert.True(IsLevelEnabled(LevelError), "Error level must be enabled")
	assert.True(IsLevelEnabled(LevelFatal), "Fatal level must be enabled")

	SetLevel(LevelError)
	assert.False(IsLevelEnabled(LevelVerbose), "Verbose level must not be enabled")
	assert.False(IsLevelEnabled(LevelDebug), "Debug level must not be enabled")
	assert.False(IsLevelEnabled(LevelInfo), "Info level must not be enabled")
	assert.False(IsLevelEnabled(LevelWarn), "Warn level must not be enabled")
	assert.True(IsLevelEnabled(LevelError), "Error level must be enabled")
	assert.True(IsLevelEnabled(LevelFatal), "Fatal level must be enabled")

	SetLevel(LevelFatal)
	assert.False(IsLevelEnabled(LevelVerbose), "Verbose level must not be enabled")
	assert.False(IsLevelEnabled(LevelDebug), "Debug level must not be enabled")
	assert.False(IsLevelEnabled(LevelInfo), "Info level must not be enabled")
	assert.False(IsLevelEnabled(LevelWarn), "Warn level must not be enabled")
	assert.False(IsLevelEnabled(LevelError), "Error level must not be enabled")
	assert.True(IsLevelEnabled(LevelFatal), "Fatal level must be enabled")
}
