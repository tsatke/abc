package abc

type RuntimeInformation struct{}

func (r *RuntimeInformation) File(depth int) string {
	return "" // calling file
}

func (r *RuntimeInformation) Line(depth int) int {
	return -1 // calling line number
}

func (r *RuntimeInformation) Function(depth int) string {
	return "" // calling package
}
