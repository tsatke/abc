package abc

import (
	"bytes"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// =================
// == DO NOT EDIT ==
// =================

func TestCustomPatternLogger_Stackops_Line(t *testing.T) {
	buf := &bytes.Buffer{}

	check := func(expectation string) {
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

	logger := &CustomPatternLogger{
		clock:   &mockClock{},
		lvl:     LevelVerbose,
		out:     buf,
		pattern: `{{.Line}}`,
	}

	// actual test flow
	logger.Verbose("")        // this is line 39
	check("39")               //
	logger.Verbosef("%v", "") // this is line 41
	check("41")               //
	logger.Debug("")          // ...
	check("43")
	logger.Debugf("%v", "")
	check("45")
	logger.Info("")
	check("47")
	logger.Infof("%v", "")
	check("49")
	logger.Warn("")
	check("51")
	logger.Warnf("%v", "")
	check("53")
	logger.Error("")
	check("55")
	logger.Errorf("%v", "")
	check("57")
	logger.Fatal("")
	check("59")
	logger.Fatalf("%v", "")
	check("61")
	logger.Print(LevelInfo, "")
	check("63")
	logger.Printf(LevelInfo, "%v", "")
	check("65")
}
