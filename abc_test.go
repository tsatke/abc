package abc

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
	if _, ok := Root().(*SimpleLogger); !ok {
		t.Error("Expected root logger to be of type *SimpleLogger, but was not.")
	}
}

func TestSetRoot(t *testing.T) {
	temp := Root()      // save original root logger
	defer SetRoot(temp) // cleanup

	// no output during tests
	l := &SimpleLogger{
		clock: &mockClock{},
		lvl:   LevelInfo,
		out:   ioutil.Discard,
	}
	SetRoot(l)

	Info("foo") // should write into ioutil.Discard

	buf := &bytes.Buffer{}

	logger := &SimpleLogger{
		clock: &mockClock{},
		lvl:   LevelInfo,
		out:   buf,
	}
	SetRoot(logger) // set new root logger

	Info("abc")

	expected := "0001-01-01 00:00:00.000 [INFO] - abc\n"
	if buf.String() != expected {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, buf.String(), false)

		t.Fail()
		t.Logf(`Expected "%v"`, expected)
		t.Logf(`Received "%v"`, buf.String())
		t.Logf("Diff: %v", dmp.DiffPrettyText(diffs))
	}
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
	_, ok = logger.clock.(*realClock)
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
	_, ok = logger.clock.(*realClock)
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
	_, ok = logger.clock.(*realClock)
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
