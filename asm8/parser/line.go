package parser

import (
	"e8vm.net/leaf/asm8/ast"
	"e8vm.net/leaf/asm8/lexer/tt"
)

func (p *Parser) parseLine() *ast.Line {
	p.push("line")
	defer p.pop()

	ret := new(ast.Line)
	err := func() *ast.Line {
		p.skipUntil(tt.Semi)
		return ret
	}

	if p.accept(tt.Semi) {
		// empty line
		return ret
	}

	if !p.accept(tt.Ident) {
		return err()
	}

	id := p.last()

	if p.ahead(tt.Colon) {
		// a label
		p.extend("label")
		assert(p.expect(tt.Colon))
		p.pop()

		ret.Label = &ast.Label{
			NameToken: id,
			Name:      id.Lit,
		}

		if p.ahead(tt.Semi) {
			// empty line label
			assert(p.expect(tt.Semi))
			return ret
		}

		// read the ident again
		if !p.expect(tt.Ident) {
			return err()
		}

		id = p.last()
	}

	inst := new(ast.Inst)
	ret.Inst = inst

	inst.OpToken = id
	inst.Op = id.Lit

	var valid bool
	inst.Args, valid = p.parseArgs()

	if !valid {
		return err()
	}

	if !p.expect(tt.Semi) {
		return err()
	}

	return ret
}

func (p *Parser) parseArgs() ([]*ast.Arg, bool) {
	var ret []*ast.Arg

	if p.ahead(tt.Semi) {
		return nil, true
	}

	for len(ret) < 5 {
		arg := p.parseArg()
		if arg == nil {
			return ret, false
		}

		ret = append(ret, arg)

		if p.ahead(tt.Semi) {
			return ret, true
		}

		if !p.expect(tt.Comma) {
			return ret, false
		}
	}

	// too many args
	p.err("too many op args in a line")
	return ret, false
}

func (p *Parser) parseArg() *ast.Arg {
	p.push("arg")
	defer p.pop()
	ret := new(ast.Arg)

	if p.accept(tt.Dollar) {
		if !p.expect(tt.Int) {
			return nil
		}

		ret.Reg = p.last()
		return ret
	}

	if p.accept(tt.Int) {
		ret.Im = p.last()
	} else if p.accept(tt.Ident) {
		ret.Sym = p.last()
	} else if !p.ahead(tt.Lparen) {
		p.expecting("instruction arg")
		return nil
	}

	// at this point, if the lparen is missing
	// then it must be a bare int or ident
	// which is okay
	if p.accept(tt.Lparen) {
		if !p.expect(tt.Dollar) {
			return nil
		}
		if !p.expect(tt.Int) {
			return nil
		}

		ret.AddrReg = p.last()

		if !p.expect(tt.Rparen) {
			return nil
		}
	}

	return ret
}
