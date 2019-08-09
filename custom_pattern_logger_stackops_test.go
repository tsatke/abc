package abc

import (
	"bytes"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func TestCustomPatternLogger_Stackops_File(t *testing.T) {
	expectation := "custom_pattern_logger_stackops_test.go custom_pattern_logger_stackops_test.go"

	buf := &bytes.Buffer{}

	logger := &CustomPatternLogger{
		clk:     &mockClock{},
		lvl:     LevelVerbose,
		out:     buf,
		pattern: `{{.File}} {{.Filef "short"}}`,
	}

	check := func() {
		defer buf.Reset()

		if buf.String() != expectation {
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(expectation, buf.String(), false)

			t.Fail()
			t.Logf(`Expected "%v"`, expectation)
			t.Logf(`Received "%v"`, buf.String())
			t.Logf("Diff: %v", dmp.DiffPrettyText(diffs))
		}
	}

	// actual test flow

	logger.Verbose("")
	check()
	logger.Verbosef("%v", "")
	check()
	logger.Debug("")
	check()
	logger.Debugf("%v", "")
	check()
	logger.Info("")
	check()
	logger.Infof("%v", "")
	check()
	logger.Warn("")
	check()
	logger.Warnf("%v", "")
	check()
	logger.Error("")
	check()
	logger.Errorf("%v", "")
	check()
	logger.Fatal("")
	check()
	logger.Fatalf("%v", "")
	check()
	logger.Print(LevelInfo, "")
	check()
	logger.Printf(LevelInfo, "%v", "")
	check()
}

func TestCustomPatternLogger_Stackops_Function(t *testing.T) {
	expectation := "abc.TestCustomPatternLogger_Stackops_Function TestCustomPatternLogger_Stackops_Function github.com/TimSatke/abc.TestCustomPatternLogger_Stackops_Function abc.TestCustomPatternLogger_Stackops_Function"

	buf := &bytes.Buffer{}

	logger := &CustomPatternLogger{
		clk:     &mockClock{},
		lvl:     LevelVerbose,
		out:     buf,
		pattern: `{{.Function}} {{.Functionf "short"}} {{.Functionf "full"}} {{.Functionf "package"}}`,
	}

	check := func() {
		defer buf.Reset()

		if buf.String() != expectation {
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(expectation, buf.String(), false)

			t.Fail()
			t.Logf(`Expected "%v"`, expectation)
			t.Logf(`Received "%v"`, buf.String())
			t.Logf("Diff: %v", dmp.DiffPrettyText(diffs))
		}
	}

	// actual test flow

	logger.Verbose("")
	check()
	logger.Verbosef("%v", "")
	check()
	logger.Debug("")
	check()
	logger.Debugf("%v", "")
	check()
	logger.Info("")
	check()
	logger.Infof("%v", "")
	check()
	logger.Warn("")
	check()
	logger.Warnf("%v", "")
	check()
	logger.Error("")
	check()
	logger.Errorf("%v", "")
	check()
	logger.Fatal("")
	check()
	logger.Fatalf("%v", "")
	check()
	logger.Print(LevelInfo, "")
	check()
	logger.Printf(LevelInfo, "%v", "")
	check()
}
