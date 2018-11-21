package abc

import "time"

type mockClock struct{}

func (mockClock) Now() time.Time                         { return time.Time{} }
func (mockClock) After(d time.Duration) <-chan time.Time { return time.After(d) }
