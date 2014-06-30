package tests

import (
	"testing"

	"e8vm.net/leaf/script"
)

func testOut(t *testing.T, src string, out string) {
	run := new(script.Run)
	run.Filename = "test.l"
	run.Source = []byte(src)
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
}
