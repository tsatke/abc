package abc

import (
	"sync"
)

var (
	mu   sync.Mutex
	root Logger
)

func init() {
	SetRoot(NewSimpleLogger())
}

func Root() Logger {
	return root
}

func SetRoot(lg Logger) {
	mu.Lock()
	defer mu.Unlock()
	root = lg
}

func NewSimpleLogger() Logger {
	return &SimpleLogger{
		lvl: LevelInfo,
	}
}

// "Implementing" abc.Logger

func Print(lvl LogLevel, v ...interface{}) {
	if root.IsLevelEnabled(lvl) {
		root.Print(lvl, v)
	}
}

func Printf(lvl LogLevel, format string, v ...interface{}) {
	if root.IsLevelEnabled(lvl) {
		root.Printf(lvl, format, v)
	}
}
