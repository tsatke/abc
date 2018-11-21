package abc

import (
	"io/ioutil"
	"testing"
)

func BenchmarkSimpleLogger_Printf(b *testing.B) {
	logger := &SimpleLogger{
		clock: &mockClock{},
		lvl:   LevelVerbose,
		out:   ioutil.Discard,
	}
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Printf(LevelVerbose, "formatted: %v", "some input")
		}
	})
}
