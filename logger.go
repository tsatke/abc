package abc

import "io"

type LogLevel uint8

const (
	LevelVerbose LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

func (l LogLevel) String() string {
	switch l {
	case LevelVerbose,
		LevelDebug:
		return "DEBG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERR"
	case LevelFatal:
		return "FATAL"
	}
	return ""
}

type Logger interface {
	Print(LogLevel, ...interface{})
	Printf(LogLevel, string, ...interface{})

	Inspect(interface{})
	Verbose(...interface{})
	Verbosef(string, ...interface{})

	Debug(...interface{})
	Debugf(string, ...interface{})

	Info(...interface{})
	Infof(string, ...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})

	Error(...interface{})
	Errorf(string, ...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})

	Out() io.Writer
	SetOut(io.Writer)

	SetLevel(LogLevel)
	IsLevelEnabled(LogLevel) bool
}
