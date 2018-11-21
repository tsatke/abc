package abc

import "io"

type WriterLogger interface {
	Logger

	Out() io.Writer
	SetOut(io.Writer)
}
