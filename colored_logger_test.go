package abc

import (
	"bytes"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestColoredLogger_Printf(t *testing.T) {
	buf := &bytes.Buffer{}

	l := &SimpleLogger{
		clock: &mockClock{},
		lvl:   LevelDebug,
		out:   buf,
	}

	logger := &ColoredLogger{
		wrapped: l,
	}

	type args struct {
		lvl    LogLevel
		format string
		v      []interface{}
	}
	tests := []struct {
		name     string
		s        *ColoredLogger
		args     args
		expected string
	}{
		{
			"Print single string with higher log level",
			logger,
			args{
				LevelInfo,
				"%v",
				[]interface{}{"abc"},
			},
			string(ColorGreen) + "0001-01-01 00:00:00.000 [INFO] - abc\n" + string(ColorReset),
		},
		{
			"Print single string with same log level",
			logger,
			args{
				LevelDebug,
				"%v",
				[]interface{}{"abc"},
			},
			string(ColorNone) + "0001-01-01 00:00:00.000 [DEBG] - abc\n" + string(ColorReset),
		},
		{
			"Print single string with higher log level (WARN)",
			logger,
			args{
				LevelWarn,
				"%v",
				[]interface{}{"abc"},
			},
			string(ColorYellow) + "0001-01-01 00:00:00.000 [WARN] - abc\n" + string(ColorReset),
		},
		{
			"Print single string with higher log level (ERR)",
			logger,
			args{
				LevelError,
				"%v",
				[]interface{}{"abc"},
			},
			string(ColorRed) + "0001-01-01 00:00:00.000 [ERR ] - abc\n" + string(ColorReset),
		},
		{
			"Print single string with higher log level (FATAL)",
			logger,
			args{
				LevelFatal,
				"%v",
				[]interface{}{"abc"},
			},
			string(ColorRed) + "0001-01-01 00:00:00.000 [FATAL] - abc\n" + string(ColorReset),
		},
		{
			"Output suppressed due to log level",
			logger,
			args{
				LevelVerbose,
				"%v",
				[]interface{}{"abc"},
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer buf.Reset()

			tt.s.Printf(tt.args.lvl, tt.args.format, tt.args.v...)
			if buf.String() != tt.expected {
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(tt.expected, buf.String(), false)

				t.Fail()
				t.Logf(`Expected "%v"`, tt.expected)
				t.Logf(`Received "%v"`, buf.String())
				t.Logf("Diff: %v", dmp.DiffPrettyText(diffs))
			}
		})
	}
}

func TestColoredLogger_Print(t *testing.T) {
	buf := &bytes.Buffer{}

	l := &SimpleLogger{
		clock: &mockClock{},
		lvl:   LevelDebug,
		out:   buf,
	}

	logger := &ColoredLogger{
		wrapped: l,
	}

	type args struct {
		lvl LogLevel
		v   []interface{}
	}
	tests := []struct {
		name     string
		s        *ColoredLogger
		args     args
		expected string
	}{
		{
			"Print single string with higher log level",
			logger,
			args{
				LevelInfo,
				[]interface{}{"abc"},
			},
			string(ColorGreen) + "0001-01-01 00:00:00.000 [INFO] - abc\n" + string(ColorReset),
		},
		{
			"Print single string with same log level",
			logger,
			args{
				LevelDebug,
				[]interface{}{"abc"},
			},
			string(ColorNone) + "0001-01-01 00:00:00.000 [DEBG] - abc\n" + string(ColorReset),
		},
		{
			"Print single string with higher log level (WARN)",
			logger,
			args{
				LevelWarn,
				[]interface{}{"abc"},
			},
			string(ColorYellow) + "0001-01-01 00:00:00.000 [WARN] - abc\n" + string(ColorReset),
		},
		{
			"Print single string with higher log level (ERR)",
			logger,
			args{
				LevelError,
				[]interface{}{"abc"},
			},
			string(ColorRed) + "0001-01-01 00:00:00.000 [ERR ] - abc\n" + string(ColorReset),
		},
		{
			"Print single string with higher log level (FATAL)",
			logger,
			args{
				LevelFatal,
				[]interface{}{"abc"},
			},
			string(ColorRed) + "0001-01-01 00:00:00.000 [FATAL] - abc\n" + string(ColorReset),
		},
		{
			"Output suppressed due to log level",
			logger,
			args{
				LevelVerbose,
				[]interface{}{"abc"},
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer buf.Reset()

			tt.s.Print(tt.args.lvl, tt.args.v...)
			if buf.String() != tt.expected {
				dmp := diffmatchpatch.New()
				diffs := dmp.DiffMain(tt.expected, buf.String(), false)

				t.Fail()
				t.Logf(`Expected "%v"`, tt.expected)
				t.Logf(`Received "%v"`, buf.String())
				t.Logf("Diff: %v", dmp.DiffPrettyText(diffs))
			}
		})
	}
}
