package parser

import (
	"io"

	"e8vm.net/leaf/lexer"
	"e8vm.net/leaf/lexer/tt"
	"e8vm.net/util/tok"
)

type scanner struct {
	lexer *lexer.Lexer
	cur   *tok.Token
	last  *tok.Token

	errors []error

	*tracker
}

func newScanner(in io.Reader, filename string) *scanner {
	ret := new(scanner)
	ret.lexer = lexer.New(in, filename)
	ret.lexer.ErrorFunc = func(e error) {
		ret.errors = append(ret.errors, e)
	}

	ret.tracker = new(tracker)
	ret.shift()

	return ret
}

func ttOf(t *tok.Token) tt.T {
	return t.Type.(tt.T)
}

// reads in the next token
// return false if the current token is already end-of-file
func (self *scanner) shift() bool {
	if self.cur != nil && ttOf(self.cur) == tt.EOF {
		return false
	}

	for self.lexer.Scan() {
		self.last = self.cur

		if self.last != nil && ttOf(self.last) != tt.Comment {
			self.tracker.add(self.last) // record in tracker
		}

		self.cur = self.lexer.Token()
		if ttOf(self.cur) != tt.Comment {
			return true
		}
	}

	panic("should never reach here")
}

func (self *scanner) ahead(tok tt.T) bool {
	return ttOf(self.cur) == tok
}

func (self *scanner) accept(tok tt.T) bool {
	if tok == tt.EOF {
		panic("cannot accept EOF")
	}

	if self.ahead(tok) {
		return self.shift()
	}
	return false
}

func (self *scanner) eof() bool {
	return self.ahead(tt.EOF)
}

func (self *scanner) skipUntil(t tt.T) []*tok.Token {
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
