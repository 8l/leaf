package parser

import (
	"e8vm.net/leaf/leaf/ast"
	"e8vm.net/leaf/leaf/lexer/tt"
)

func (p *Parser) parseTopDecl() ast.Node {
	if p.ahead(tt.Func) {
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

	// TODO: parse args and signature
	if !p.expect(tt.Lparen) {
		return err()
	}

	if !p.expect(tt.Rparen) {
		return err()
	}

	if !p.ahead(tt.Lbrace) {
		p.expect(tt.Lbrace)
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
	p.skipUntil(tt.Semi)
}
