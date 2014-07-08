package parser

import (
	"errors"
	"fmt"

	"e8vm.net/leaf/asm8/lexer/tt"
	"e8vm.net/leaf/tools/comperr"
	"e8vm.net/leaf/tools/tok"
	"e8vm.net/leaf/tools/tracker"
)

// inherits helper functions from the scanner
func (p *Parser) last() *tok.Token       { return p.s.Last() }
func (p *Parser) cur() *tok.Token        { return p.s.Cur() }
func (p *Parser) pop() tracker.Node      { return p.s.Pop() }
func (p *Parser) push(s string)          { p.s.Push(s) }
func (p *Parser) extend(s string)        { p.s.Extend(s) }
func (p *Parser) ahead(t tok.Type) bool  { return p.s.Ahead(t) }
func (p *Parser) accept(t tok.Type) bool { return p.s.Accept(t) }
func (p *Parser) shift() bool            { return p.s.Shift() }
func (p *Parser) eof() bool              { return p.s.EOF() }

func (p *Parser) skipUntil(t tok.Type) []*tok.Token {
	return p.s.SkipUntil(t)
}

// Errors returns the a list of errors encountered.
// It returns the scanner errors first if there is any.
// Otherwise, it will return the parsing error if any.
func (p *Parser) Errors() []error {
	errors := p.s.Errors()
	if errors != nil {
		return errors
	}

	return p.errors
}

// err appends an error onto the error list
func (p *Parser) err(s string) {
	e := comperr.New(p.cur(), errors.New(s))
	p.errors = append(p.errors, e)
}

// expect helpers
func (p *Parser) expect(tok tt.T) bool {
	if p.ahead(tok) {
		assert(p.shift())
		return true
	}

	p.err(fmt.Sprintf("expect %s, got %s", tok, p.cur().Type))

	p.shift() // make progress anyway
	return false
}

/*
func (p *Parser) expectSemi() bool {
	if p.ahead(tt.Rparen) || p.ahead(tt.Rbrace) {
		return true
	}

	return p.expect(tt.Semi)
}
*/

func (p *Parser) aheadSemi() bool {
	if p.ahead(tt.Rparen) || p.ahead(tt.Rbrace) {
		return true
	}

	return p.ahead(tt.Semi)
}

func (p *Parser) expecting(s string) {
	p.err(fmt.Sprintf("expect %s, got %s", s, p.cur().Type))
}

// asserting
func assert(cond bool) {
	if !cond {
		panic("bug")
	}
}
