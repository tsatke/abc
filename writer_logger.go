package abc

import "io"

// WriterLogger is an interface that embeds abc.Logger.
// It describes loggers that write their output to
// an io.Writer, and provides methods for retrieving
// and changing that writer.
type WriterLogger interface {
	Logger

	// Out returns the writer to which the output
	// is printed.
	Out() io.Writer
	// SetOut changes the writer to which the output
	// is printed.
	SetOut(io.Writer)
}
