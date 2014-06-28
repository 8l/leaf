package codegen

import (
	"strconv"

	"e8vm.net/leaf/ir"
	"e8vm.net/leaf/ir/types"
	"e8vm.net/leaf/lexer"
	"e8vm.net/leaf/lexer/token"
	"e8vm.net/leaf/parser/ast"
)

func (self *Gen) genExpr(code *ir.Code, expr ast.Node) *obj {
	switch expr := expr.(type) {
	default:
		panic("bug or todo")
	case *ast.CallExpr:
		return self.genCall(code, expr)
	case *ast.Operand:
		return self.genOperand(code, expr)
	}
}

func (self *Gen) genCall(code *ir.Code, call *ast.CallExpr) *obj {
	f := self.genExpr(code, call.Func) // evaluate the function first
	if f == nil {
		return nil
	}

	ft, isFunc := f.t.(*types.Func)
	if !isFunc {
		self.errorf(call.Token, "calling on a non-function")
		return nil
	}

	if len(call.Args) != len(ft.Args) {
		self.errorf(call.Token, "wrong number of arguments")
		return nil
	}

	var args []*obj
	for i, arg := range call.Args {
		o := self.genExpr(code, arg)
		if o == nil {
			return nil
		}
		if !types.Equals(o.t, ft.Args[i]) {
			self.errorf(call.Token, "wrong argument type")
			return nil
		}
		args = append(args, o)
	}

	// TODO: push the ret first
	ret := voidObj

	for _, o := range args {
		code.Push(o.o) // now we can push the stuff for call
	}

	code.Call(f.o)

	var pops []ir.Obj
	for _, o := range args {
		pops = append(pops, o.o)
	}
	code.Pop(pops...)

	return ret
}

func (self *Gen) genOperand(code *ir.Code, op *ast.Operand) *obj {
	tok := op.Token

	switch tok.Token {
	default:
		panic("bug or todo")

	case token.Int:
		// TODO: use real large integers
		i, e := strconv.ParseInt(tok.Lit, 0, 64)
		if e != nil {
			self.errore(tok, e)
			return nil
		}
		return &obj{ir.ConstNum(i), types.ConstNum}
	case token.Char:
		c, e := unquoteChar(tok.Lit)
		if e != nil {
			self.errore(tok, e)
			return nil
		}
		return &obj{ir.ConstInt(int64(c), types.Int8), types.Int8}
	case token.Ident:
		return self.genIdent(code, tok)
	}
}

func (self *Gen) genIdent(code *ir.Code, tok *lexer.Token) *obj {
	assert(tok.Token == token.Ident)

	o, t := code.Query(tok.Lit)
	if o == nil {
		self.errorf(tok, "%q undefined", tok.Lit)
		return nil
	}

	return &obj{o, t}
}
