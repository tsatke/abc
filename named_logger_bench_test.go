package abc

import (
	"testing"
)

func BenchmarkNamedLogger_Printf(b *testing.B) {
	logger := &NamedLogger{
		clock: &mockClock{},
		lvl:   LevelVerbose,
		out:   &MockWriter{},
		name:  "MyLogger",
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Printf(LevelVerbose, "formatted: %v", "some input")
		}
	})
}
