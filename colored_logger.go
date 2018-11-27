package abc

import (
	"io"
	"sync"
)

type color []byte

// Available colors (16 bit ANSI color codes)
var (
	ColorGray   = color("\u001b[30;1m")
	ColorGreen  = color("\u001b[32;1m")
	ColorNone   = ColorReset
	ColorRed    = color("\u001b[31m")
	ColorReset  = color("\u001b[0m")
	ColorYellow = color("\u001b[33m")
)

// ColoredLogger is a wrapper for any WriterLogger.
// Depending on the level that should be printed, this wrapper
// will prepend an ANSI-color code to the wrapped loggers
// output writer and will then call the respective output method.
// Please notice that this is not add a decorator for
// a given logger, but a wrapper, which must be used
// for colors to show up.
//
// The color codes are written to the wrapped loggers output
// writer.
// An output call will trigger three writes on the wrapped loggers
// output writer.
//
// 1. color code
//
// 2. output from wrapped logger
//
// 3. color code reset
//
// The second call will not originate from the wrapper,
// as, after the color code was written, Print (or respectively Printf)
// will be called on the wrapped logger.
type ColoredLogger struct {
	wrappedLock sync.Mutex
	wrapped     WriterLogger
}

func (s *ColoredLogger) printWithColor(clr color, lvl LogLevel, v ...interface{}) {
	s.wrappedLock.Lock()
	defer s.wrappedLock.Unlock()

	if s.IsLevelEnabled(lvl) {
		s.wrapped.Out().Write(clr)
		s.wrapped.Print(lvl, v...)
		s.wrapped.Out().Write(ColorReset)
	}
}

func (s *ColoredLogger) printfWithColor(clr color, lvl LogLevel, format string, v ...interface{}) {
	s.wrappedLock.Lock()
	defer s.wrappedLock.Unlock()

	if s.IsLevelEnabled(lvl) {
		s.wrapped.Out().Write(clr)
		s.wrapped.Printf(lvl, format, v...)
		s.wrapped.Out().Write(ColorReset)
	}
}

// Print delegates the values with the given log level to the wrapped
// logger while writing ansi color codes to the wrapped loggers output writer.
func (s *ColoredLogger) Print(lvl LogLevel, v ...interface{}) {
	color := s.getColorForLevel(lvl)
	s.printWithColor(color, lvl, v...)
}

// Printf delegates the format string and values with the given log level to the wrapped
// logger while writing ansi color codes to the wrapped loggers output writer.
func (s *ColoredLogger) Printf(lvl LogLevel, format string, v ...interface{}) {
	color := s.getColorForLevel(lvl)
	s.printfWithColor(color, lvl, format, v...)
}

func (s *ColoredLogger) getColorForLevel(lvl LogLevel) color {
	switch lvl {
	case LevelVerbose:
		return ColorGray
	case LevelDebug:
		return ColorNone
	case LevelInfo:
		return ColorGreen
	case LevelWarn:
		return ColorYellow
	case LevelError:
		return ColorRed
	case LevelFatal:
		return ColorRed
	default:
		return ColorNone
	}
}

// Verbose delegates the given values to the wrapped logger
// while writing ansi color codes to the wrapped loggers output writer.
func (s *ColoredLogger) Verbose(v ...interface{}) {
	s.Print(LevelVerbose, v...)
}

// Verbosef delegates the given values and format string to the wrapped logger
// while writing ansi color codes to the wrapped loggers output writer.
func (s *ColoredLogger) Verbosef(format string, v ...interface{}) {
	s.Printf(LevelVerbose, format, v...)
}

// Debug delegates the given values to the wrapped logger
// while writing ansi color codes to the wrapped loggers output writer.
func (s *ColoredLogger) Debug(v ...interface{}) {
	s.Print(LevelDebug, v...)
}

// Debugf delegates the given values to the wrapped logger
// while writing ansi color codes to the wrapped loggers output writer.
func (s *ColoredLogger) Debugf(format string, v ...interface{}) {
	s.Printf(LevelDebug, format, v...)
}

// Info delegates the given values to the wrapped logger
// while writing ansi color codes to the wrapped loggers output writer.
func (s *ColoredLogger) Info(v ...interface{}) {
	s.Print(LevelInfo, v...)
}

// Infof delegates the given values to the wrapped logger
// while writing ansi color codes to the wrapped loggers output writer.
func (s *ColoredLogger) Infof(format string, v ...interface{}) {
	s.Printf(LevelInfo, format, v...)
}

// Warn delegates the given values to the wrapped logger
// while writing ansi color codes to the wrapped loggers output writer.
func (s *ColoredLogger) Warn(v ...interface{}) {
	s.Print(LevelWarn, v...)
}

// Warnf delegates the given values to the wrapped logger
// while writing ansi color codes to the wrapped loggers output writer.
func (s *ColoredLogger) Warnf(format string, v ...interface{}) {
	s.Printf(LevelWarn, format, v...)
}

// Error delegates the given values to the wrapped logger
// while writing ansi color codes to the wrapped loggers output writer.
func (s *ColoredLogger) Error(v ...interface{}) {
	s.Print(LevelError, v...)
}

// Errorf delegates the given values to the wrapped logger
// while writing ansi color codes to the wrapped loggers output writer.
func (s *ColoredLogger) Errorf(format string, v ...interface{}) {
	s.Printf(LevelError, format, v...)
}

// Fatal delegates the given values to the wrapped logger
// while writing ansi color codes to the wrapped loggers output writer.
//
// THIS WRAPPER DOES NOT TERMINATE THE APPLICATION AFTER A FATAL CALL.
func (s *ColoredLogger) Fatal(v ...interface{}) {
	s.Print(LevelFatal, v...)
}

// Fatalf delegates the given values to the wrapped logger
// while writing ansi color codes to the wrapped loggers output writer.
//
// THIS WRAPPER DOES NOT TERMINATE THE APPLICATION AFTER A FATAL CALL.
func (s *ColoredLogger) Fatalf(format string, v ...interface{}) {
	s.Printf(LevelFatal, format, v...)
}

// SetLevel delegates the given log level to the wrapped logger.
func (s *ColoredLogger) SetLevel(lvl LogLevel) {
	s.wrapped.SetLevel(lvl)
}

// IsLevelEnabled delegates to the wrapped loggers IsLevelEnabled method.
func (s *ColoredLogger) IsLevelEnabled(lvl LogLevel) bool {
	return s.wrapped.IsLevelEnabled(lvl)
}

// Out returns the writer of the wrapped logger.
func (s *ColoredLogger) Out() io.Writer {
	return s.wrapped.Out()
}

// SetOut sets a new writer for the wrapped logger.
func (s *ColoredLogger) SetOut(out io.Writer) {
	s.wrapped.SetOut(out)
}
