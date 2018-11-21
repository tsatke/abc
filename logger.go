package abc

// Logger describes objects that can log messages.
// It can differentiate between several log levels,
// and provides methods for each one.
type Logger interface {
	// Print prints the given values with the given log level,
	// if and only if the given log level is higher than or
	// equal to the one of this logger.
	Print(LogLevel, ...interface{})
	// Printf formats and prints the given values with the given log level,
	// if and only if the given log level is higher than or
	// equal to the one of this logger.
	Printf(LogLevel, string, ...interface{})

	// Inspect prints detailed information about the given value.
	// Very verbose, may be slow.
	// NOT RECOMMENDED FOR PRODUCTION USE.
	Inspect(interface{})
	// Verbose prints the given values with log level DEBG,
	// if and only if this logger has the verbose log level enabled.
	Verbose(...interface{})
	// Verbosef formats and prints the given values with log level DEBG,
	// if and only if this logger has the verbose log level enabled.
	Verbosef(string, ...interface{})

	// Debug prints the given values with log level DEBG.
	Debug(...interface{})
	// Debugf formats and prints the given values with log level DEBG.
	Debugf(string, ...interface{})

	// Info prints the given values with log level INFO.
	Info(...interface{})
	// Infof formats and prints the given values with log level INFO.
	Infof(string, ...interface{})

	// Warn prints the given values with log level WARN.
	Warn(...interface{})
	// Warnf formats and prints the given values with log level WARN.
	Warnf(string, ...interface{})

	// Error prints the given values with log level ERR.
	Error(...interface{})
	// Errorf formats and prints the given values with log level ERR.
	Errorf(string, ...interface{})

	// Fatal prints the given values with log level FATAL.
	// IT DOES NOT TERMINATE THE APPLICATION.
	Fatal(...interface{})
	// Fatalf formats and prints the given values with log level FATAL.
	// IT DOES NOT TERMINATE THE APPLICATION.
	Fatalf(string, ...interface{})

	// SetLevel changes the log level of this logger.
	SetLevel(LogLevel)
	// IsLevelEnabled returns true if and only if this logger would print
	// messages with the given log level.
	// False otherwise.
	IsLevelEnabled(LogLevel) bool
}
