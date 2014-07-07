// Package lexer parses an input file into tokens.
package lexer

import (
	"fmt"
	"io"

	"e8vm.net/leaf/lexer/tt"
	"e8vm.net/util/comperr"
	"e8vm.net/util/runes"
	"e8vm.net/util/scanner"
	"e8vm.net/util/tok"
)

type Lexer struct {
	s   *scanner.Scanner
	buf *tok.Token

	illegal    bool   // illegal encountered
	insertSemi bool   // if treat end line as whitespace
	eof        bool   // end of file returned
	err        error  // first error encountered
	filename   string // filename for printing error

	ErrorFunc func(e error)
}

// Creates a new lexer
func New(in io.Reader, filename string) *Lexer {
	ret := new(Lexer)
	ret.s = scanner.New(in)
	ret.buf = new(tok.Token)
	ret.buf.File = filename
	ret.filename = filename

	return ret
}

func (self *Lexer) wrapError(e error) error {
	ret := new(comperr.Error)
	ret.Err = e
	ret.Line, ret.Col = self.s.Pos()
	ret.File = self.filename

	return ret
}

func (self *Lexer) report(e error) {
	if e == nil {
		return
	}

	e = self.wrapError(e)

	if self.err == nil {
		self.err = e
	}

	if self.ErrorFunc != nil {
		self.ErrorFunc(e)
	}
}

// Reports a lex error
func (self *Lexer) failf(f string, args ...interface{}) {
	self.report(fmt.Errorf(f, args...))
}

func (self *Lexer) skipWhites() {
	if self.insertSemi {
		self.s.SkipAnys(" \t\r")
	} else {
		self.s.SkipAnys(" \t\r\n")
	}
}

func (self *Lexer) _scanNumber(dotLed bool) (lit string, t tt.T) {
	s := self.s

	if !dotLed {
		if s.Scan('0') {
			if s.Scan('x') || self.s.Scan('X') {
				if s.ScanHexDigits() == 0 {
					return s.Accept(), tt.Illegal
				}
			} else if s.ScanOctDigit() {
				s.ScanOctDigits()
				return s.Accept(), tt.Int
			}

			if s.Peek() != '.' {
				return s.Accept(), tt.Int
			}
		}

		s.ScanDigits()

		if s.ScanAny("eE") {
			s.ScanAny("-+")
			if s.ScanDigits() == 0 {
				return s.Accept(), tt.Illegal
			}
			return s.Accept(), tt.Float
		}

		if !s.Scan('.') {
			return s.Accept(), tt.Int
		}

		s.ScanDigits()
	} else {
		if s.ScanDigits() == 0 {
			return s.Accept(), tt.Illegal
		}
	}

	if s.ScanAny("eE") {
		s.ScanAny("-+")
		if s.ScanDigits() == 0 {
			return s.Accept(), tt.Illegal
		}
	}

	return s.Accept(), tt.Float
}

func (self *Lexer) scanNumber(dotLed bool) (lit string, t tt.T) {
	lit, t = self._scanNumber(dotLed)
	if t == tt.Illegal {
		self.failf("invalid number")
		t = tt.Int
	}

	return
}

func (self *Lexer) scanEscape(q rune) {
	s := self.s

	if s.ScanAny("abfnrtv\\") {
		return
	}
	if s.Scan(q) {
		return
	}

	if s.Scan('x') {
		if !(s.ScanHexDigit() && s.ScanHexDigit()) {
			self.failf("invalid hex escape")
		}
		return
	}

	if s.ScanOctDigit() {
		if !(s.ScanOctDigit() && s.ScanOctDigit()) {
			self.failf("invalid octal escape")
		}
		return
	}

	self.failf("unknown escape char %q", s.Peek())
	s.Next()

	return
}

func (self *Lexer) scanChar() string {
	s := self.s
	n := 0
	for !s.Scan('\'') {
		if s.Peek() == '\n' || s.Closed() {
			self.failf("char not terminated")
			break
		}

		if s.Scan('\\') {
			self.scanEscape('\'')
		} else {
			s.Next()
		}
		n++
	}

	if n != 1 {
		self.failf("illegal char")
	}

	return s.Accept()
}

func (self *Lexer) scanString() string {
	s := self.s

	for !s.Scan('"') {
		if s.Peek() == '\n' || s.Closed() {
			self.failf("string not terminated")
			break
		}

		if s.Scan('\\') {
			self.scanEscape('"')
		} else {
			s.Next()
		}
	}

	return s.Accept()
}

func (self *Lexer) scanRawString() string {
	s := self.s

	for !s.Scan('`') {
		if s.Closed() {
			self.failf("raw string not terminated")
			break
		}
		s.Next()
	}
	return s.Accept()
}

func (self *Lexer) scanComment() string {
	s := self.s

	if s.Scan('*') {
		for {
			if s.Scan('*') {
				if s.Scan('/') {
					return s.Accept()
				}
				continue
			}

			if s.Closed() {
				self.failf("incomplete block comment")
				return s.Accept()
			}
			s.Next()
		}
	}

	if s.Scan('/') {
		for {
			if s.Peek() == '\n' || s.Closed() {
				return s.Accept()
			}
			s.Next()
		}
	}

	panic("bug")
}

func (self *Lexer) Err() error {
	return self.err
}

func (self *Lexer) ScanErr() error {
	return self.s.Err()
}

var insertSemiTokens = []tt.T{
	tt.Ident,
	tt.Int,
	tt.Float,
	tt.Break,
	tt.Continue,
	tt.Fallthrough,
	tt.Return,
	tt.Char,
	tt.String,
	tt.Rparen,
	tt.Rbrack,
	tt.Rbrace,
	tt.Inc,
	tt.Dec,
}

var insertSemiTokenMap = func() map[tt.T]bool {
	ret := make(map[tt.T]bool)
	for _, t := range insertSemiTokens {
		ret[t] = true
	}
	return ret
}()

func (self *Lexer) savePos() {
	self.buf.Line, self.buf.Col = self.s.Pos()
}

func (self *Lexer) token(t tt.T, lit string) *tok.Token {
	self.buf.Type = t
	self.buf.Lit = lit
	return self.buf
}

// Returns if the scanner has anything to return
func (self *Lexer) Scan() bool { return !self.eof }

// Returns the next token.
// t is the token code, p is the position code,
// and lit is the string literal.
// Returns token.EOF in t for the last token.
func (self *Lexer) Token() *tok.Token {
	ret := self.scanToken()
	t := ret.Type.(tt.T)
	if t != tt.Illegal {
		self.insertSemi = insertSemiTokenMap[t]
	}

	return ret.Clone()
}

func (self *Lexer) scanToken() *tok.Token {
	if self.eof {
		// once it reached eof, it will repeatedly return EOF
		self.savePos()
		return self.token(tt.EOF, "")
	}

	self.skipWhites()
	self.savePos()

	if self.s.Closed() {
		if self.insertSemi {
			self.insertSemi = false
			return self.token(tt.Semi, ";")
		}
		self.eof = true

		self.report(self.s.Err())

		return self.token(tt.EOF, "")
	}

	s := self.s
	r := s.Peek()

	switch {
	case runes.IsLetter(r):
		s.ScanIdent()
		lit := s.Accept()
		t := tt.FromIdent(lit)
		return self.token(t, lit)
	case runes.IsDigit(r):
		lit, t := self.scanNumber(false)
		return self.token(t, lit)
	case r == '\'':
		s.Next()
		lit := self.scanChar()
		return self.token(tt.Char, lit)
	case r == '"':
		s.Next()
		lit := self.scanString()
		return self.token(tt.String, lit)
	case r == '`':
		s.Next()
		lit := self.scanRawString()
		return self.token(tt.String, lit)
	}

	s.Next() // at this time, we will always make some progress

	if r == '.' && runes.IsDigit(s.Peek()) {
		lit, t := self.scanNumber(true)
		return self.token(t, lit)
	} else if r == '/' {
		r2 := s.Peek()
		if r2 == '/' || r2 == '*' {
			s := self.scanComment()
			return self.token(tt.Comment, s)
		}
	}

	t := self.scanOperator(r)
	lit := s.Accept()
	if t == tt.Semi {
		lit = ";"
	}

	return self.token(t, lit)
}
