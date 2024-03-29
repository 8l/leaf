package parser

import (
	"e8vm.net/leaf/leaf/ast"
	"e8vm.net/leaf/leaf/lexer/tt"
	"e8vm.net/leaf/tools/tok"
)

func (p *Parser) parseStmt() ast.Node {
	switch {
	case p.ahead(tok.EOF):
		p.err("unexpected EOF")
		return nil

	case p.cur().Type.(tt.T).IsLiteral():
		fallthrough
	case p.ahead(tt.Ident) || p.ahead(tt.Lparen):
		return p.parseExprStmt()

	case p.ahead(tt.Semi):
		return p.parseEmptyStmt()

	case p.ahead(tt.Lbrace):
		return p.parseBlock()

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
		p.skipUntil(tt.Semi)
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

	if !p.expect(tt.Lbrace) {
		return ret
	}

	for !p.ahead(tt.Rbrace) {
		if p.ahead(tok.EOF) {
			break
		}
		stmt := p.parseStmt()
		if stmt != nil {
			ret.Stmts = append(ret.Stmts, stmt)
		}
	}

	p.expect(tt.Rbrace)
	return ret
}

func (p *Parser) parseErrorStmt() {
	p.push("error-stmt")
	defer p.pop()

	p.expecting("statement")
	p.skipUntil(tt.Semi)
}
