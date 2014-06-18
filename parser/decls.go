package parser

import (
	"e8vm.net/leaf/ast"
	t "e8vm.net/leaf/token"
)

func (p *Parser) parseTopDecl() ast.Node {
	if p.ahead(t.Func) {
		return p.parseFunc()
	}

	p.parseErrorDecl()
	return nil
}

func (p *Parser) parseFunc() *ast.Func {
	p.push("func-decl")
	defer p.pop()

	ret := new(ast.Func)
	err := func() *ast.Func {
		p.skipUntil(t.Semi)
		return ret
	}

	if !p.expect(t.Func) {
		return err()
	}

	if !p.expect(t.Ident) {
		return err()
	}

	ret.Name = p.last.Lit

	// TODO: parse args and signature
	if !p.expect(t.Lparen) {
		return err()
	}

	if !p.expect(t.Rparen) {
		return err()
	}

	if !p.ahead(t.Lbrace) {
		p.expect(t.Lbrace)
		return err()
	}

	ret.Block = p.parseBlock()

	if !p.expectSemi() {
		return err()
	}

	return ret
}

func (p *Parser) parseErrorDecl() {
	p.push("error-decl")
	defer p.pop()

	p.expecting("declaration")
	p.skipUntil(t.Semi)
}
