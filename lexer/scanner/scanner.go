// Package scanner defines a general text scanner.
package scanner

import (
	"bufio"
	"bytes"
	"io"

	"e8vm.net/leaf/lexer/runes"
)

const (
	MaxRunePerLine = 250   // 8 bit
	MaxLinePerFile = 50000 // 16 bit
)

type Scanner struct {
	reader *bufio.Reader

	r rune

	tail *pos
	head *pos

	buf    *bytes.Buffer
	closed bool
	err    error
}

func New(in io.Reader) *Scanner {
	ret := new(Scanner)
	ret.reader = bufio.NewReader(in)
	ret.head = newPos()
	ret.tail = newPos()
	ret.buf = new(bytes.Buffer)

	ret.Next() // get ready for reading

	return ret
}

func (self *Scanner) shutdown(e error) {
	self.r = rune(-1)
	self.closed = true
	self.err = e
}

func (self *Scanner) Closed() bool {
	return self.closed
}

// Returns the scanning error.
func (self *Scanner) Err() error { return self.err }

// Returns the current position of the tail pointer.
func (self *Scanner) Pos() (int, int) {
	return self.tail.lineNo, self.tail.lineOffset
}

// Increase head pointer by one rune.
func (self *Scanner) Next() rune {
	if self.closed {
		return rune(-1)
	}

	self.buf.WriteRune(self.r)

	var rsize int
	var e error
	self.r, rsize, e = self.reader.ReadRune()
	if e == io.EOF {
		self.shutdown(nil)
	} else if e != nil {
		self.shutdown(e)
	}

	if self.r == '\n' {
		self.head.NewLine()
	} else {
		self.head.lineOffset += rsize
	}

	return self.r
}

// Returns the rune at the head pointer.
func (self *Scanner) Peek() rune {
	return self.r
}

// Increase the head pointer by one if the rune at it is r.
// Returns true if the head moved.
func (self *Scanner) Scan(r rune) bool {
	if self.r == r {
		self.Next()
		return true
	}
	return false
}

// Increase the head pointer by one if f returns true for the rune pointing
// Returns true if the head moved.
func (self *Scanner) ScanFunc(f func(r rune) bool) bool {
	if f(self.r) {
		self.Next()
		return true
	}
	return false
}

// Increase the head pointer until f returns false for the rune pointing
// Returns the number of runes that the head reads.
func (self *Scanner) ScanFuncs(f func(r rune) bool) int {
	ret := 0
	for self.ScanFunc(f) {
		ret++
	}

	return ret
}

// Increase the head pointer until the rune is not a digit
// Returns the number of runes that the head reads
func (self *Scanner) ScanDigits() int {
	return self.ScanFuncs(runes.IsDigit)
}

// Increase the head pointer by one if the rune is a hex digit
// Returns true if the head moved.
func (self *Scanner) ScanHexDigit() bool {
	return self.ScanFunc(runes.IsHexDigit)
}

// Increase the head pointer by one if the rune is a octal digit
// Returns true if the head moved.
func (self *Scanner) ScanOctDigit() bool {
	return self.ScanFunc(runes.IsOctDigit)
}

// Increase the head pointer until the rune is not a hex digit
// Returns the number of runes that the head reads
func (self *Scanner) ScanHexDigits() int {
	return self.ScanFuncs(runes.IsHexDigit)
}

// Increase the head pointer until the rune is not a octal digit
// Returns the number of runes that the head reads
func (self *Scanner) ScanOctDigits() int {
	return self.ScanFuncs(runes.IsOctDigit)
}

// Increase the head pointer by one if the rune is a letter or '_'.
// Returns true if the head pointer moved.
func (self *Scanner) ScanLetter() bool {
	return self.ScanFunc(runes.IsLetter)
}

// Incrase the head pointer by one if the rune is a digit
// Retunrs true if the head pointer moved.
func (self *Scanner) ScanDigit() bool {
	return self.ScanFunc(runes.IsDigit)
}

// Increase the head pointer until the rune is either a digit or a letter
// (it takes '_' as a letter).
// Returns the number of runes that the head reads.
func (self *Scanner) ScanIdent() int {
	ret := 0
	for runes.IsDigit(self.r) || runes.IsLetter(self.r) {
		ret++
		self.Next()
	}
	return ret
}

// Increase the head pointer if the rune is in string s.
// Returns true if the head moved.
func (self *Scanner) ScanAny(s string) bool {
	for _, r := range s {
		if r == self.r {
			self.Scan(rune(r))
			return true
		}
	}

	return false
}

// Increase the head pointer until the rune is not in string s.
// Returns the number of runes that the head reads.
func (self *Scanner) ScanAnys(s string) int {
	ret := 0
	for self.ScanAny(s) {
		ret++
	}
	return ret
}

func (self *Scanner) SyncTail() {
	self.buf.Reset()
	self.tail.SyncTo(self.head)
}

// Returns the string captured by the tail and the head, and
// sync the tail to the head.
func (self *Scanner) Accept() string {
	ret := self.buf.String()
	self.SyncTail()

	return ret
}

// Increase the head pointer until the rune is not in string s.
// Also sync the tail to the head.
// Returns the number of runes that the head reads.
func (self *Scanner) SkipAnys(s string) int {
	ret := self.ScanAnys(s)
	self.SyncTail()
	return ret
}
