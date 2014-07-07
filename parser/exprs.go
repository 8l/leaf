package parser

import (
	"e8vm.net/leaf/ast"
	"e8vm.net/leaf/lexer/tt"
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

	if p.ahead(tt.Lparen) {
		return p.parseCall(x)
	}
	return x
}

func (p *Parser) parseCall(f ast.Node) ast.Node {
	p.extend("call-expr")
	defer p.pop()

	assert(p.expect(tt.Lparen)) // otherwise, why you are here?

	ret := new(ast.CallExpr)
	ret.Func = f
	ret.Token = p.last()

	for p.until(tt.Rparen) {
		arg := p.parseExpr()
		ret.Args = append(ret.Args, arg)

		if p.ahead(tt.Rparen) {
			continue
		}

		if !p.expect(tt.Comma) {
			continue
		}
	}

	if !p.expect(tt.Rparen) {
		p.skipUntil(tt.Rparen)
	}

	return ret
}

// parse identifiers, literals, and paren'ed expressions
// prefix with t.Ident, t.Literals, and t.Lparen
func (p *Parser) parseOperand() ast.Node {
	p.push("operand")
	defer p.pop()

	switch {
	case p.accept(tt.Ident):
		ret := new(ast.Operand)
		ret.Token = p.last()
		return ret
	case p.cur().Type.(tt.T).IsLiteral():
		p.shift()
		ret := new(ast.Operand)
		ret.Token = p.last()
		return ret
	case p.ahead(tt.Lparen):
		return p.parseParenExpr()
	}

	p.expecting("operand")
	return nil
}

func (p *Parser) parseParenExpr() ast.Node {
	if !p.expect(tt.Lparen) {
		return nil
	}

	ret := p.parseExpr()

	if !p.expect(tt.Rparen) {
		p.skipUntil(tt.Rparen)
	}

	return ret
}
