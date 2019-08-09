package abc

import (
	"testing"
)

func BenchmarkSimpleLogger_Printf(b *testing.B) {
	logger := &SimpleLogger{
		clk: &mockClock{},
		lvl: LevelVerbose,
		out: &MockWriter{},
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Printf(LevelVerbose, "formatted: %v", "some input")
		}
	})
}
