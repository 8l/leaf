// Package lexer parses an input file into tokens
package lexer

import (
	"io"

	"e8vm.net/leaf/asm8/lexer/tt"

	"e8vm.net/leaf/tools/runes"
	"e8vm.net/leaf/tools/scanner"
	"e8vm.net/leaf/tools/tok"
)

// Lexer parses an input stream into a stream of tokens
type Lexer struct {
	s        *scanner.Scanner
	buf      *tok.Token
	filename string // filename for printing error

	illegal    bool  // illegal encountered
	insertSemi bool  // if treat end line as whitespaces
	eof        bool  // end of file returned
	err        error // first error encountered

	// onError is called when an error is encountered
	onError func(e error)
}

// New creates a new lexer for an input stream.
func New(in io.Reader, filename string) *Lexer {
	ret := new(Lexer)
	ret.s = scanner.New(in, filename)
	ret.buf = new(tok.Token)
	ret.buf.File = filename // always use the same filename
	ret.filename = filename

	return ret
}

// OnError sets the error callback function for lexing errors.
func (lx *Lexer) OnError(f func(e error)) {
	lx.onError = f
}

func (lx *Lexer) report(e error) {
	if e == nil {
		return
	}

	if lx.err == nil {
		lx.err = e
	}
	if lx.onError != nil {
		lx.onError(e)
	}
}

func (lx *Lexer) reportf(f string, args ...interface{}) {
	lx.report(lx.s.Errorf(f, args...))
}

func (lx *Lexer) skipWhites() {
	if lx.insertSemi {
		lx.s.SkipAnys(" \t\r")
	} else {
		lx.s.SkipAnys(" \t\r\n")
	}
}

func (lx *Lexer) scanNumber(dotLed bool) (lit string, t tt.T) {
	lit, ntype := scanner.ScanNumber(lx.s, dotLed)
	switch ntype {
	case scanner.NumIllegal:
		t = tt.Int
		lx.reportf("invalid number")
	case scanner.NumInt:
		t = tt.Int
	case scanner.NumFloat:
		t = tt.Float
	default:
		panic("bug")
	}
	return lit, t
}

// Err returns the first error encountered, including scanning error.
func (lx *Lexer) Err() error { return lx.err }

// ScanErr returns the first error encounted on scanning.
func (lx *Lexer) ScanErr() error { return lx.s.Err() }

var insertSemiTokens = []tt.T{
	tt.Ident,
	tt.Int,
	tt.Float,
	tt.Char,
	tt.String,
	tt.Rparen,
	tt.Rbrace,
	tt.Colon,
}

var insertSemiTokenMap = func() map[tt.T]bool {
	ret := make(map[tt.T]bool)
	for _, t := range insertSemiTokens {
		ret[t] = true
	}
	return ret
}()

func (lx *Lexer) savePos() { lx.buf.Line, lx.buf.Col = lx.s.Pos() }

func (lx *Lexer) token(t tt.T, lit string) *tok.Token {
	lx.buf.Type = t
	lx.buf.Lit = lit
	return lx.buf
}

func (lx *Lexer) eofToken() *tok.Token {
	lx.buf.Type = tok.EOF
	lx.buf.Lit = ""
	return lx.buf
}

// Scan tests if the scanner has anything to return
func (lx *Lexer) Scan() bool { return !lx.eof }

// Token returns the next token.
func (lx *Lexer) Token() *tok.Token {
	ret := lx.scanToken()
	t, isT := ret.Type.(tt.T)
	if isT && t != tt.Illegal && t != tt.Comment {
		lx.insertSemi = insertSemiTokenMap[t]
	}

	return ret.Clone()
}

func (lx *Lexer) scanToken() *tok.Token {
	if lx.eof {
		panic("no more")
	}

	lx.skipWhites()
	lx.savePos()

	if lx.s.Closed() {
		if lx.insertSemi {
			lx.insertSemi = false
			return lx.token(tt.Semi, ";")
		}
		lx.eof = true
		lx.report(lx.s.Err())

		return lx.eofToken()
	}

	s := lx.s
	r := s.Peek()

	switch {
	case runes.IsLetter(r):
		s.ScanIdent()
		lit := s.Accept()
		t := tt.FromIdent(lit)
		return lx.token(t, lit)
	case r == '-' || r == '+':
		s.Next() // eat the sign
		if s.Scan('.') {
			lit, t := lx.scanNumber(true)
			return lx.token(t, lit)
		} else if runes.IsDigit(s.Peek()) {
			lit, t := lx.scanNumber(false)
			return lx.token(t, lit)
		} else {
			lx.reportf("a sign must lead a number")
			lit := s.Accept()
			return lx.token(tt.Illegal, lit)
		}
	case runes.IsDigit(r):
		lit, t := lx.scanNumber(false)
		return lx.token(t, lit)
	case r == '\'': // start of a char
		s.Next()
		lit, e := scanner.ScanChar(lx.s)
		lx.report(e)
		return lx.token(tt.Char, lit)
	case r == '"': // start of a string
		s.Next()
		lit, e := scanner.ScanString(lx.s)
		lx.report(e)
		return lx.token(tt.String, lit)
	case r == '`': // start of a raw string
		s.Next()
		lit, e := scanner.ScanRawString(lx.s)
		lx.report(e)
		return lx.token(tt.String, lit)
	}

	s.Next() // always make some progress here

	if r == '.' && runes.IsDigit(s.Peek()) {
		lit, t := lx.scanNumber(true)
		return lx.token(t, lit)
	}

	if r == '/' { // start of comment
		next := s.Peek()
		if next == '/' || next == '*' {
			s, e := scanner.ScanComment(lx.s)
			lx.report(e)
			return lx.token(tt.Comment, s)
		}
	}

	t := lx.scanOperator(r)
	lit := s.Accept()
	if t == tt.Semi {
		lit = ";"
	}

	return lx.token(t, lit)
}
