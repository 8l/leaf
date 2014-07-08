package parser

import (
	"e8vm.net/leaf/asm8/ast"
	"e8vm.net/leaf/asm8/lexer/tt"
)

func (p *Parser) parseDecls() ast.Decl {
	switch {
	case p.ahead(tt.Func):
		return p.parseFunc()
	case p.ahead(tt.Var):
		panic("todo")
		return p.parseVar()
	case p.ahead(tt.Const):
		panic("todo")
		return p.parseVar()
	}

	p.parseErrorDecl()
	return nil
}

func (p *Parser) parseErrorDecl() {
	p.push("error-decl")
	defer p.pop()
	p.expecting("declaration")
	p.skipUntil(tt.Semi)
}

// parses a function declaration
func (p *Parser) parseFunc() *ast.Func {
	p.push("func")
	defer p.pop()

	ret := new(ast.Func)
	err := func() *ast.Func {
		p.skipUntil(tt.Semi)
		return ret
	}

	if !p.expect(tt.Func) {
		return err()
	}

	if !p.expect(tt.Ident) {
		return err()
	}

	ret.NameToken = p.last()
	ret.Name = ret.NameToken.Lit

	if !p.ahead(tt.Lbrace) {
		p.expect(tt.Lbrace)
		return err()
	}

	ret.Block = p.parseBlock()
	if !p.expect(tt.Semi) {
		return err()
	}

	return ret
}

func (p *Parser) parseBlock() *ast.Block {
	p.push("block")
	defer p.pop()

	ret := new(ast.Block)

	if !p.expect(tt.Lbrace) {
		return ret
	}

	for !p.ahead(tt.Rbrace) {
		if p.eof() {
			break
		}

		line := p.parseLine()
		if line != nil {
			ret.Lines = append(ret.Lines, line)
		}
	}

	p.expect(tt.Rbrace)
	return ret
}

func (p *Parser) parseVar() *ast.Var {
	p.push("var")
	defer p.pop()

	panic("todo")
}

func (p *Parser) parseConst() *ast.Const {
	p.push("const")
	defer p.pop()

	panic("todo")
}
