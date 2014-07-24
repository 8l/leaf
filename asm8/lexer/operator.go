package lexer

import (
	"e8vm.net/leaf/asm8/lexer/tt"
)

func (lx *Lexer) scanOperator(r rune) tt.T {
	switch r {
	case '\n':
		lx.insertSemi = false
		return tt.Semi
	case '(':
		return tt.Lparen
	case ')':
		return tt.Rparen
	case '{':
		return tt.Lbrace
	case '}':
		return tt.Rbrace
	case '[':
		return tt.Lbrack
	case ']':
		return tt.Rbrack
	case '=':
		return tt.Assign
	case '$':
		return tt.Dollar
	case ':':
		return tt.Colon
	case ';':
		return tt.Semi
	case ',':
		return tt.Comma
	}

	if !lx.illegal {
		lx.illegal = true
		lx.reportf("illegal character")
	}
	return tt.Illegal
}
