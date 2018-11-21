package abc

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

	SetLevel(LogLevel)
	IsLevelEnabled(LogLevel) bool
}
