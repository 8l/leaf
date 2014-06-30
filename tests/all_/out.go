package all_

import (
	"testing"

	"e8vm.net/leaf/script"
)

func match(expect []byte, got []byte) bool {
	if len(expect) != len(got) {
		return false
	}

	for i := range expect {
		if expect[i] != got[i] {
			return false
		}
	}

	return true
}

func testOut(t *testing.T, src []byte, out []byte) {
	run := new(script.Run)
	run.Filename = "test.l"
	run.Source = src
	run.Run()

	if len(run.Errors) > 0 {
		for _, e := range run.Errors {
			t.Error(e)
		}
	}

	if !run.RIP() {
		t.Errorf("vm exit with half value %d", run.HaltValue)
		if run.AddrError {
			t.Errorf("vm has address error")
		}
	}

	if !match(out, run.Output) {
		t.Errorf("output does not match")
		t.Logf("  expect: %q", out)
		t.Logf("  got: %q", run.Output)
	}
}
