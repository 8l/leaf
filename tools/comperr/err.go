// Package comperr defines a compiler error with filename and the error
// position.
package comperr

import (
	"fmt"

	"e8vm.net/leaf/tools/tok"
)

// Error is a compiler error with filename and the error position
type Error struct {
	// Err is the error message.
	Err error

	// File is the filename of the error.
	File string

	// Line is the line number.
	Line int

	// Col is the column number.
	Col int
}

// Error returns the line description of the compiler error.
func (e *Error) Error() string {
	prefix := ""
	if e.File != "" {
		prefix = e.File + ":"
	}
	return fmt.Sprintf("%s%d:%d: %s",
		prefix, e.Line, e.Col, e.Err,
	)
}

// New creates a compiler error at a given token.
func New(t *tok.Token, e error) *Error {
	ret := new(Error)
	ret.File = t.File
	ret.Line = t.Line
	ret.Col = t.Col
	ret.Err = e

	return ret
}
