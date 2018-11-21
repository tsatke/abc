package abc

// LogLevel is an alias and represents a log level.
type LogLevel uint8

// Available log levels.
//
//	LevelVerbose LogLevel = iota
//	LevelDebug
//	LevelInfo
//	LevelWarn
//	LevelError
//	LevelFatal
const (
	LevelVerbose LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

// String returns a string representation of the level
// that can be used in the log output.
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
