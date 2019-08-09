package abc

import (
	"testing"
)

func BenchmarkCustomPatternLogger_Printf(b *testing.B) {
	logger := &CustomPatternLogger{
		clk:     &mockClock{},
		lvl:     LevelVerbose,
		out:     &MockWriter{},
		pattern: "{{.Timestamp}} [{{.Level}}] - {{.Message}}\n",
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Printf(LevelVerbose, "formatted: %v", "some input")
		}
	})
}
func BenchmarkCustomPatternLogger_Printf_Stack_Ops(b *testing.B) {
	logger := &CustomPatternLogger{
		clk:     &mockClock{},
		lvl:     LevelVerbose,
		out:     &MockWriter{},
		pattern: "{{.Timestamp}} {{.File}}:{{.Line}} {{.Function}} [{{.Level}}] - {{.Message}}\n",
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Printf(LevelVerbose, "formatted: %v", "some input")
		}
	})
}
