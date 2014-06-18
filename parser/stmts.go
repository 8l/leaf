package parser

import (
	"e8vm.net/leaf/ast"
	t "e8vm.net/leaf/token"
)

func (p *Parser) parseStmt() ast.Node {
	switch {
	case p.cur.Token.IsLiteral():
		fallthrough
	case p.ahead(t.Ident) || p.ahead(t.Lparen):
		return p.parseExprStmt()

	case p.ahead(t.Semi):
		return p.parseEmptyStmt()

	case p.ahead(t.Lbrace):
		return p.parseBlock()

	case p.ahead(t.EOF):
		p.err("unexpected EOF")
		return nil

	default:
		p.parseErrorStmt()
		return nil
	}
}

func (p *Parser) parseExprStmt() *ast.ExprStmt {
	p.push("expr-stmt")
	defer p.pop()

	ret := new(ast.ExprStmt)
	ret.Expr = p.parseExpr()

	if !p.expectSemi() {
		p.skipUntil(t.Semi)
	}

	return ret
}

func (p *Parser) parseEmptyStmt() *ast.EmptyStmt {
	p.push("empty-stmt")
	p.expectSemi()
	p.pop()

	return new(ast.EmptyStmt)
}

func (p *Parser) parseBlock() *ast.Block {
	p.push("block-stmt")
	defer p.pop()

	ret := new(ast.Block)

	if !p.expect(t.Lbrace) {
		return ret
	}

	for !p.ahead(t.Rbrace) {
		if p.ahead(t.EOF) {
			break
		}
		stmt := p.parseStmt()
		if stmt != nil {
			ret.Stmts = append(ret.Stmts, stmt)
		}
	}

	p.expect(t.Rbrace)
	return ret
}

func (p *Parser) parseErrorStmt() {
	p.push("error-stmt")
	defer p.pop()

	p.expecting("statement")
	p.skipUntil(t.Semi)
}
