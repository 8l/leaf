package tests

import (
	"testing"
)

func TestHi(t *testing.T) {
	prog := `
	func main() {
		putc('h')
		putc('i')
	}
	`
	output := "hi"

	testOut(t, prog, output)
}
