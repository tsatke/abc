package abc

import (
	"io/ioutil"
	"testing"
)

func BenchmarkCustomPatternLogger_Printf(b *testing.B) {
	logger := &CustomPatternLogger{
		clock:   &mockClock{},
		lvl:     LevelVerbose,
		out:     ioutil.Discard,
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
