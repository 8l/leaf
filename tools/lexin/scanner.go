// Package lexin provides a lexer scanner that helps
// reading a token stream and fold the stream into a tree
// using a tracker.
package lexin

import (
	"io"

	"e8vm.net/leaf/tools/tok"
	"e8vm.net/leaf/tools/tracker"
)

// Lexer defines a general interface for lexing
type Lexer interface {
	Scan() bool            // test if there is token in the stream
	Token() *tok.Token     // returns the next token and shifts the cursor
	OnError(func(e error)) // set error handling function
}

// Scanner takes a lexer
type Scanner struct {
	lexer Lexer
	cur   *tok.Token
	last  *tok.Token

	errors []error

	tracker *tracker.Tracker

	SkipFunc func(*tok.Token) bool
}

// NewScanner creates a new scanner that wraps
// the given lexer.
func NewScanner(lexer Lexer, skip func(*tok.Token) bool) *Scanner {
	ret := new(Scanner)
	ret.lexer = lexer
	ret.lexer.OnError(func(e error) {
		ret.errors = append(ret.errors, e)
	})

	ret.tracker = tracker.New()
	ret.SkipFunc = skip
	ret.Shift()

	return ret
}

func (s *Scanner) ignore(t *tok.Token) bool {
	if t.Is(tok.EOF) {
		return false
	}

	if s.SkipFunc == nil {
		return false
	}

	return s.SkipFunc(t)
}

// Shift reads in the next token and returns true.
// It returns false when there is nothing to shift (meats EOF).
func (s *Scanner) Shift() bool {
	if s.cur != nil && s.cur.Is(tok.EOF) {
		return false // already shifts to the end
	}

	if !s.lexer.Scan() {
		panic("lexer not ending with EOF")
	}

	cur := s.lexer.Token() 
	for s.ignore(cur) {
		cur = s.lexer.Token()
	}

	s.last = s.cur
	if s.last != nil && !s.ignore(s.last) {
		s.tracker.Add(s.last)
	}
	s.cur = cur
	return true
}

// Ahead tests is the current token is of the specified type.
func (s *Scanner) Ahead(t tok.Type) bool {
	return s.cur.Is(t)
}

// Accept shifts when the current token is of the specified type.
// It returns true when it shifted.
// It will panic if the type is EOF.
func (s *Scanner) Accept(t tok.Type) bool {
	if t == tok.EOF {
		panic("cannot accept EOF")
	}

	if s.Ahead(t) {
		return s.Shift()
	}

	return false
}

// EOF returns true if the current token is EOF already.
func (s *Scanner) EOF() bool {
	return s.Ahead(tok.EOF)
}

// SkipUtil shifts until a particular type of token is shifted
// (it also shifts the last one).
// It returns all tokens that it shipped.
func (s *Scanner) SkipUntil(t tok.Type) []*tok.Token {
	var skipped []*tok.Token

	for !s.Ahead(t) {
		skipped = append(skipped, s.cur)
		if !s.Shift() {
			return skipped
		}
	}

	// and we shift the last one
	skipped = append(skipped, s.cur)
	s.Shift()

	return skipped
}

// Push pushes a new level on to the tracker's stack.
func (s *Scanner) Push(name string) {
	s.tracker.Push(name)
}

// Pop pops a level from the tracker's stack.
func (s *Scanner) Pop() tracker.Node {
	return s.tracker.Pop()
}

// Extend extends the last token on the tracker into a new level
// on the tracker's stack.
func (s *Scanner) Extend(name string) {
	s.tracker.Extend(name)
}

// Errors returns all the scanning errors.
func (s *Scanner) Errors() []error {
	return s.errors
}

// Last returns the last token.
func (s *Scanner) Last() *tok.Token {
	return s.last
}

// Cur returns the current scanning token.
func (s *Scanner) Cur() *tok.Token {
	return s.cur
}

// PrintTrack prints the token track tree.
func (s *Scanner) PrintTrack(out io.Writer) {
	s.tracker.Print(out)
}
