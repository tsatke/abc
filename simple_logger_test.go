package abc

import (
	"bytes"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestSimpleLogger_Printf(t *testing.T) {
	buf := &bytes.Buffer{}

	logger := &SimpleLogger{
		clock: &mockClock{},
		lvl:   LevelDebug,
		out:   buf,
	}

	type args struct {
		lvl    LogLevel
		format string
		v      []interface{}
	}
	tests := []struct {
		name     string
		s        *SimpleLogger
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
			"0001-01-01 00:00:00.000 [INFO] - abc\n",
		},
		{
			"Print single string with same log level",
			logger,
			args{
				LevelDebug,
				"%v",
				[]interface{}{"abc"},
			},
			"0001-01-01 00:00:00.000 [DEBG] - abc\n",
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

func TestSimpleLogger_Print(t *testing.T) {
	buf := &bytes.Buffer{}

	logger := &SimpleLogger{
		clock: &mockClock{},
		lvl:   LevelDebug,
		out:   buf,
	}

	type args struct {
		lvl LogLevel
		v   []interface{}
	}
	tests := []struct {
		name     string
		s        *SimpleLogger
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
			"0001-01-01 00:00:00.000 [INFO] - abc\n",
		},
		{
			"Print single string with same log level",
			logger,
			args{
				LevelDebug,
				[]interface{}{"abc"},
			},
			"0001-01-01 00:00:00.000 [DEBG] - abc\n",
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
		{
			"Print slice of strings",
			logger,
			args{
				LevelDebug,
				[]interface{}{[]string{"a", "b", "c"}},
			},
			"0001-01-01 00:00:00.000 [DEBG] - [a b c]\n",
		},
		{
			"Print struct",
			logger,
			args{
				LevelDebug,
				[]interface{}{struct {
					a string
					b float64
				}{
					a: "abc",
					b: -0.1,
				}},
			},
			"0001-01-01 00:00:00.000 [DEBG] - {abc -0.1}\n",
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

func TestSimpleLogger_All_Outputs(t *testing.T) {
	expectations := []string{
		"0001-01-01 00:00:00.000 [DEBG] - verbose: abc\n",
		"0001-01-01 00:00:00.000 [DEBG] - verbose: fmt: abc\n",
		"0001-01-01 00:00:00.000 [DEBG] - abc\n",
		"0001-01-01 00:00:00.000 [DEBG] - fmt: abc\n",
		"0001-01-01 00:00:00.000 [INFO] - abc\n",
		"0001-01-01 00:00:00.000 [INFO] - fmt: abc\n",
		"0001-01-01 00:00:00.000 [WARN] - abc\n",
		"0001-01-01 00:00:00.000 [WARN] - fmt: abc\n",
		"0001-01-01 00:00:00.000 [ERR ] - abc\n",
		"0001-01-01 00:00:00.000 [ERR ] - fmt: abc\n",
		"0001-01-01 00:00:00.000 [FATAL] - abc\n",
		"0001-01-01 00:00:00.000 [FATAL] - fmt: abc\n",
	}
	cnt := 0

	buf := &bytes.Buffer{}

	logger := &SimpleLogger{
		clock: &mockClock{},
		lvl:   LevelVerbose,
		out:   buf,
	}

	check := func() {
		defer func() {
			cnt++
		}()
		defer buf.Reset()

		if cnt >= len(expectations) {
			panic("No more expectations")
		}

		if buf.String() != expectations[cnt] {
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(expectations[cnt], buf.String(), false)

			t.Fail()
			t.Logf(`Expected "%v"`, expectations[cnt])
			t.Logf(`Received "%v"`, buf.String())
			t.Logf("Diff: %v", dmp.DiffPrettyText(diffs))
		}
	}

	// actual test flow

	logger.Verbose("verbose: abc")
	check()
	logger.Verbosef("verbose: fmt: %v", "abc")
	check()
	logger.Debug("abc")
	check()
	logger.Debugf("fmt: %v", "abc")
	check()
	logger.Info("abc")
	check()
	logger.Infof("fmt: %v", "abc")
	check()
	logger.Warn("abc")
	check()
	logger.Warnf("fmt: %v", "abc")
	check()
	logger.Error("abc")
	check()
	logger.Errorf("fmt: %v", "abc")
	check()
	logger.Fatal("abc")
	check()
	logger.Fatalf("fmt: %v", "abc")
	check()
}
