// Package lexer parses an input file into tokens.
package lexer

import (
	"io"

	"e8vm.net/leaf/lexer/tt"
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
	ret.s = scanner.New(in, filename)
	ret.buf = new(tok.Token)
	ret.buf.File = filename
	ret.filename = filename

	return ret
}

func (lx *Lexer) report(e error) {
	if e == nil {
		return
	}
	if lx.err == nil {
		lx.err = e
	}

	if lx.ErrorFunc != nil {
		lx.ErrorFunc(e)
	}
}

func (lx *Lexer) reportf(f string, args ...interface{}) {
	e := lx.s.Errorf(f, args...)
	lx.report(e)
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

func (lx *Lexer) Err() error {
	return lx.err
}

func (lx *Lexer) ScanErr() error {
	return lx.s.Err()
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

func (lx *Lexer) savePos() {
	lx.buf.Line, lx.buf.Col = lx.s.Pos()
}

func (lx *Lexer) token(t tt.T, lit string) *tok.Token {
	lx.buf.Type = t
	lx.buf.Lit = lit
	return lx.buf
}

// Returns if the scanner has anything to return
func (lx *Lexer) Scan() bool { return !lx.eof }

// Returns the next token.
// t is the token code, p is the position code,
// and lit is the string literal.
// Returns token.EOF in t for the last token.
func (lx *Lexer) Token() *tok.Token {
	ret := lx.scanToken()
	t := ret.Type.(tt.T)
	if t != tt.Illegal {
		lx.insertSemi = insertSemiTokenMap[t]
	}

	return ret.Clone()
}

func (lx *Lexer) scanToken() *tok.Token {
	if lx.eof {
		// once it reached eof, it will repeatedly return EOF
		lx.savePos()
		return lx.token(tt.EOF, "")
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

		return lx.token(tt.EOF, "")
	}

	s := lx.s
	r := s.Peek()

	switch {
	case runes.IsLetter(r):
		s.ScanIdent()
		lit := s.Accept()
		t := tt.FromIdent(lit)
		return lx.token(t, lit)
	case runes.IsDigit(r):
		lit, t := lx.scanNumber(false)
		return lx.token(t, lit)
	case r == '\'':
		s.Next()
		lit, e := scanner.ScanChar(lx.s)
		lx.report(e)
		return lx.token(tt.Char, lit)
	case r == '"':
		s.Next()
		lit, e := scanner.ScanString(lx.s)
		lx.report(e)
		return lx.token(tt.String, lit)
	case r == '`':
		s.Next()
		lit, e := scanner.ScanRawString(lx.s)
		lx.report(e)
		return lx.token(tt.String, lit)
	}

	s.Next() // at this time, we will always make some progress

	if r == '.' && runes.IsDigit(s.Peek()) {
		lit, t := lx.scanNumber(true)
		return lx.token(t, lit)
	} else if r == '/' {
		r2 := s.Peek()
		if r2 == '/' || r2 == '*' {
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
