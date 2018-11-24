package abc

import (
	"log"
	"testing"
)

func BenchmarkStdLogger(b *testing.B) {
	logger := log.New(&MockWriter{}, "", log.Ldate|log.Ltime|log.Lmicroseconds)
	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Printf("[%v] formatted: %v", "INFO", "some input")
		}
	})
}

func BenchmarkStdLogger_Stackops(b *testing.B) {
	logger := log.New(&MockWriter{}, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Printf("[%v] formatted: %v", "INFO", "some input")
		}
	})
}
