// Package scanner defines a general text scanner.
package scanner

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	"e8vm.net/leaf/tools/comperr"
	"e8vm.net/leaf/tools/runes"
)

// Scanner scans over a io.Reader, often used as a helper for lexing.
type Scanner struct {
	reader *bufio.Reader

	r rune

	tail *pos
	head *pos

	buf    *bytes.Buffer
	closed bool
	err    error

	filename string
}

// New returns a new scanner.
func New(in io.Reader, filename string) *Scanner {
	ret := new(Scanner)
	ret.reader = bufio.NewReader(in)
	ret.head = newPos()
	ret.tail = newPos()
	ret.buf = new(bytes.Buffer)
	ret.filename = filename

	ret.Next() // get ready for reading

	return ret
}

func (s *Scanner) shutdown(e error) {
	s.r = rune(-1)
	s.closed = true
	s.err = e
}

// Closed tests if the scanner is closed.
func (s *Scanner) Closed() bool {
	return s.closed
}

// Err returns the scanning error.
func (s *Scanner) Err() error { return s.err }

// Pos returns the current position of the tail pointer.
func (s *Scanner) Pos() (int, int) {
	return s.tail.lineNo, s.tail.lineOffset
}

// Errorf creates a compiler error using fmt.Errorf
func (s *Scanner) Errorf(f string, args ...interface{}) error {
	return s.CompErr(fmt.Errorf(f, args...))
}

// CompErr wraps an error into a compiler error.
func (s *Scanner) CompErr(e error) error {
	ret := new(comperr.Error)
	ret.Err = e
	ret.Line, ret.Col = s.Pos()
	ret.File = s.filename
	return ret
}

// Next increases head pointer by one rune.
// It returns -1 when the scanner is closed.
func (s *Scanner) Next() rune {
	if s.closed {
		return rune(-1)
	}

	s.buf.WriteRune(s.r)

	var rsize int
	var e error
	s.r, rsize, e = s.reader.ReadRune()
	if e == io.EOF {
		s.shutdown(nil)
	} else if e != nil {
		s.shutdown(e)
	}

	if s.r == '\n' {
		s.head.newLine()
	} else {
		s.head.lineOffset += rsize
	}

	return s.r
}

// Peek returns the rune at the head pointer.
func (s *Scanner) Peek() rune {
	return s.r
}

// Scan increases the head pointer by one if the rune at it is r.
// It returns true if the head moved.
func (s *Scanner) Scan(r rune) bool {
	if s.r == r {
		s.Next()
		return true
	}
	return false
}

// ScanFunc increase the head pointer by one if f returns true for the rune
// pointing.  It returns true if the head moved.
func (s *Scanner) ScanFunc(f func(r rune) bool) bool {
	if f(s.r) {
		s.Next()
		return true
	}
	return false
}

// ScanFuncs increases the head pointer until f returns false for the rune
// pointing.  It returns the number of runes that the head reads.
func (s *Scanner) ScanFuncs(f func(r rune) bool) int {
	ret := 0
	for s.ScanFunc(f) {
		ret++
	}

	return ret
}

// ScanDigits increases the head pointer until the rune is not a digit.
// It returns the number of runes that the head reads
func (s *Scanner) ScanDigits() int {
	return s.ScanFuncs(runes.IsDigit)
}

// ScanHexDigit increases the head pointer by one if the rune is a hex digit.
// It returns true if the head moved.
func (s *Scanner) ScanHexDigit() bool {
	return s.ScanFunc(runes.IsHexDigit)
}

// ScanOctDigit increases the head pointer by one if the rune is a octal digit.
// It returns true if the head moved.
func (s *Scanner) ScanOctDigit() bool {
	return s.ScanFunc(runes.IsOctDigit)
}

// ScanHexDigits increases the head pointer until the rune is not a hex digit.
// It returns the number of runes that the head reads.
func (s *Scanner) ScanHexDigits() int {
	return s.ScanFuncs(runes.IsHexDigit)
}

// ScanOctDigits increases the head pointer until the rune is not a octal digit
// It returns the number of runes that the head reads
func (s *Scanner) ScanOctDigits() int {
	return s.ScanFuncs(runes.IsOctDigit)
}

// ScanLetter increases the head pointer by one if the rune is a letter or '_'.
// It returns true if the head pointer moved.
func (s *Scanner) ScanLetter() bool {
	return s.ScanFunc(runes.IsLetter)
}

// ScanDigit increases the head pointer by one if the rune is a digit It
// returns true if the head pointer moved.
func (s *Scanner) ScanDigit() bool {
	return s.ScanFunc(runes.IsDigit)
}

// ScanIdent increases the head pointer until the rune is either a digit or a
// letter (it treats '_' as a letter).  It returns the number of runes that the
// head reads.
func (s *Scanner) ScanIdent() int {
	ret := 0
	for runes.IsDigit(s.r) || runes.IsLetter(s.r) {
		ret++
		s.Next()
	}
	return ret
}

// ScanAny increases the head pointer if the rune is in string s.
// It returns true if the head moved.
func (s *Scanner) ScanAny(str string) bool {
	for _, r := range str {
		if r == s.r {
			s.Scan(rune(r))
			return true
		}
	}

	return false
}

// ScanAnys increases the head pointer until the rune is not in string s.
// It returns the number of runes that the head reads.
func (s *Scanner) ScanAnys(str string) int {
	ret := 0
	for s.ScanAny(str) {
		ret++
	}
	return ret
}

// SyncTail synchronizes the tail pointer to the head and discards the runes in
// between.
func (s *Scanner) SyncTail() {
	s.buf.Reset()
	s.tail.syncTo(s.head)
}

// Accept returns the string captured by the tail and the head, and
// synchronizes the tail to the head.
func (s *Scanner) Accept() string {
	ret := s.buf.String()
	s.SyncTail()

	return ret
}

// SkipAnys increases the head pointer until the rune is not in string s, and
// synchronizes the tail to the head.  It Returns the number of runes that the
// head reads.
func (s *Scanner) SkipAnys(str string) int {
	ret := s.ScanAnys(str)
	s.SyncTail()
	return ret
}
