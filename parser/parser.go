package parser

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"e8vm.net/leaf/ast"
	"e8vm.net/leaf/lexer"
	"e8vm.net/leaf/lexer/tt"
	"e8vm.net/util/comperr"
	"e8vm.net/util/lexin"
	"e8vm.net/util/tok"
	"e8vm.net/util/tracker"
)

type Parser struct {
	filename string
	in       io.Reader

	lex    *lexer.Lexer
	s      *lexin.Scanner
	errors []error
}

func Parse(f string) (*ast.Program, []error) {
	p, e := Open(f)
	if e != nil {
		return nil, []error{e}
	}

	return p.Parse()
}

func ParseStr(f string, prog string) (*ast.Program, []error) {
	in := strings.NewReader(prog)
	p := New(in, f)
	return p.Parse()
}

func ParseBytes(f string, prog []byte) (*ast.Program, []error) {
	in := bytes.NewBuffer(prog)
	p := New(in, f)
	return p.Parse()
}

func isComment(t *tok.Token) bool {
	return t.Is(tt.Comment)
}

func New(in io.Reader, filename string) *Parser {
	ret := new(Parser)
	ret.filename = filename
	ret.lex = lexer.New(in, filename)
	ret.s = lexin.NewScanner(ret.lex)

	ret.s.SkipFunc = isComment

	return ret
}

func Open(filename string) (*Parser, error) {
	fin, e := os.Open(filename)
	if e != nil {
		return nil, e
	}

	return New(fin, filename), nil
}

func (p *Parser) Parse() (*ast.Program, []error) {
	p.push("source-file")
	defer p.pop()

	ret := new(ast.Program)
	ret.Filename = p.filename

	for !p.s.EOF() {
		d := p.parseTopDecl()
		if d != nil {
			ret.Decls = append(ret.Decls, d)
		}
	}

	return ret, p.Errors()
}

func (p *Parser) Errors() []error {
	errors := p.s.Errors()
	if errors != nil {
		return errors
	}

	return p.errors
}

func (p *Parser) expect(tok tt.T) bool {
	if p.ahead(tok) {
		assert(p.shift())
		return true
	}

	p.err(fmt.Sprintf("expect %s, got %s", tok, p.cur().Type))

	p.shift() // make progress anyway
	return false
}

func (p *Parser) expectSemi() bool {
	if p.ahead(tt.Rparen) || p.ahead(tt.Rbrace) {
		return true
	}

	return p.expect(tt.Semi)
}

func (p *Parser) expecting(s string) {
	p.err(fmt.Sprintf("expect %s, got %s", s, p.cur().Type))
}

func (p *Parser) err(s string) {
	e := comperr.New(p.cur(), errors.New(s))
	p.errors = append(p.errors, e)
}

func (p *Parser) until(tok tt.T) bool {
	if p.s.EOF() {
		return false
	}
	if p.ahead(tok) {
		return false
	}

	return true
}

func (p *Parser) last() *tok.Token       { return p.s.Last() }
func (p *Parser) cur() *tok.Token        { return p.s.Cur() }
func (p *Parser) pop() tracker.Node      { return p.s.Pop() }
func (p *Parser) push(s string)          { p.s.Push(s) }
func (p *Parser) extend(s string)        { p.s.Extend(s) }
func (p *Parser) ahead(t tok.Type) bool  { return p.s.Ahead(t) }
func (p *Parser) accept(t tok.Type) bool { return p.s.Accept(t) }
func (p *Parser) shift() bool            { return p.s.Shift() }

func (p *Parser) skipUntil(t tok.Type) []*tok.Token {
	return p.s.SkipUntil(t)
}
