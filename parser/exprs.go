package parser

import (
	t "e8vm.net/leaf/lexer/token"
	"e8vm.net/leaf/parser/ast"
)

func (p *Parser) parseExpr() ast.Node {
	p.push("expr")
	defer p.pop()

	return p.parseBinaryExpr()
}

func (p *Parser) parseBinaryExpr() ast.Node {
	return p.parseUnaryExpr()
}

func (p *Parser) parseUnaryExpr() ast.Node {
	return p.parsePrimaryExpr()
}

func (p *Parser) parsePrimaryExpr() ast.Node {
	x := p.parseOperand()

	if p.ahead(t.Lparen) {
		return p.parseCall(x)
	}
	return x
}

func (p *Parser) parseCall(f ast.Node) ast.Node {
	p.extend("call-expr")
	defer p.pop()

	assert(p.expect(t.Lparen)) // otherwise, why you are here?

	ret := new(ast.CallExpr)
	ret.Func = f

	for p.until(t.Rparen) {
		arg := p.parseExpr()
		ret.Args = append(ret.Args, arg)

		if p.ahead(t.Rparen) {
			continue
		}

		if !p.expect(t.Comma) {
			continue
		}
	}

	if !p.expect(t.Rparen) {
		p.skipUntil(t.Rparen)
	}

	return ret
}

// parse identifiers, literals, and paren'ed expressions
// prefix with t.Ident, t.Literals, and t.Lparen
func (p *Parser) parseOperand() ast.Node {
	p.push("operand")
	defer p.pop()

	switch {
	case p.accept(t.Ident):
		ret := new(ast.Operand)
		ret.Token = p.last
		return ret
	case p.cur.Token.IsLiteral():
		p.shift()
		ret := new(ast.Operand)
		ret.Token = p.last
		return ret
	case p.ahead(t.Lparen):
		return p.parseParenExpr()
	}

	p.expecting("operand")
	return nil
}

func (p *Parser) parseParenExpr() ast.Node {
	if !p.expect(t.Lparen) {
		return nil
	}

	ret := p.parseExpr()

	if !p.expect(t.Rparen) {
		p.skipUntil(t.Rparen)
	}

	return ret
}
