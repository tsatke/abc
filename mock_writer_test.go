package abc

type MockWriter struct {
	msg string
}

func (m MockWriter) Write(b []byte) (int, error) {
	_ = string(b)
	return 0, nil
}
