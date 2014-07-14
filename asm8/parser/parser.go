// Package parser parses an assembly source file into an AST.
package parser

import (
	"io"

	"e8vm.net/leaf/asm8/ast"
	"e8vm.net/leaf/asm8/lexer"
	"e8vm.net/leaf/asm8/lexer/tt"
	"e8vm.net/leaf/tools/lexin"
	"e8vm.net/leaf/tools/tok"
)

// Parser parses an assembly source file into an AST.
type Parser struct {
	filename string
	in       io.Reader

	lx     *lexer.Lexer
	s      *lexin.Scanner
	errors []error
}

// New creates a new parser that parses a file into an AST.
func New(in io.Reader, filename string) *Parser {
	ret := new(Parser)
	ret.filename = filename
	ret.lx = lexer.New(in, filename)
	ret.s = lexin.NewScanner(ret.lx, func(t *tok.Token) bool {
		return t.Is(tt.Comment)
	})

	return ret
}

// Parse parses the file.
func (p *Parser) Parse() (*ast.Program, []error) {
	p.push("program")
	defer p.pop()

	ret := new(ast.Program)
	ret.Filename = p.filename

	for !p.s.EOF() {
		d := p.parseDecls()
		if d != nil {
			ret.Decls = append(ret.Decls, d)
		}
	}

	ret.EOFToken = p.s.Cur()

	return ret, p.Errors()
}

// PrintTree prints the token tree to an output stream.
func (p *Parser) PrintTree(out io.Writer) {
	p.s.PrintTrack(out)
}
