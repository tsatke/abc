package abc

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
