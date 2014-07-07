package parser

import (
	"io"

	"e8vm.net/leaf/lexer"
	"e8vm.net/leaf/lexer/tt"
	"e8vm.net/util/tok"
	"e8vm.net/util/tracker"
)

type scanner struct {
	lexer *lexer.Lexer
	cur   *tok.Token
	last  *tok.Token

	errors []error

	tracker *tracker.Tracker
}

func newScanner(in io.Reader, filename string) *scanner {
	ret := new(scanner)
	ret.lexer = lexer.New(in, filename)
	ret.lexer.ErrorFunc = func(e error) {
		ret.errors = append(ret.errors, e)
	}

	ret.tracker = tracker.New()
	ret.shift()

	return ret
}

func ttIs(t *tok.Token, x tok.Type) bool {
	return t.Type.Code() == x.Code()
}

// reads in the next token
// return false if the current token is already end-of-file
func (self *scanner) shift() bool {
	if self.cur != nil && ttIs(self.cur, tok.EOF) {
		return false
	}

	for self.lexer.Scan() {
		self.last = self.cur

		if self.last != nil && ttIs(self.last, tt.Comment) {
			self.add(self.last) // record in tracker
		}

		self.cur = self.lexer.Token()
		if !ttIs(self.cur, tt.Comment) {
			return true
		}
	}

	panic("should never reach here")
}

func (self *scanner) ahead(tok tok.Type) bool {
	return ttIs(self.cur, tok)
}

func (self *scanner) accept(t tok.Type) bool {
	if t == tok.EOF {
		panic("cannot accept EOF")
	}

	if self.ahead(t) {
		return self.shift()
	}
	return false
}

func (self *scanner) eof() bool {
	return self.ahead(tok.EOF)
}

func (self *scanner) skipUntil(t tok.Type) []*tok.Token {
	var skipped []*tok.Token

	for !self.ahead(t) {
		skipped = append(skipped, self.cur)
		if !self.shift() {
			return skipped
		}
	}

	// shift the last one
	skipped = append(skipped, self.cur)
	self.shift()

	return skipped
}

func (s *scanner) push(str string) {
	s.tracker.Push(str)
}

func (s *scanner) pop() tracker.Node {
	return s.tracker.Pop()
}

func (s *scanner) extend(str string) {
	s.tracker.Extend(str)
}

func (s *scanner) add(n tracker.Node) {
	s.tracker.Add(n)
}
