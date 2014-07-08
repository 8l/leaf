package lexer

import (
	"e8vm.net/leaf/leaf/lexer/tt"
)

var simpleOps = map[rune]tt.T{
	':': tt.Colon,
	',': tt.Comma,
	';': tt.Semi,
	'(': tt.Lparen,
	')': tt.Rparen,
	'[': tt.Lbrack,
	']': tt.Rbrack,
	'{': tt.Lbrace,
	'}': tt.Rbrace,
}

var eqOps = map[rune]*struct{ t, eqt tt.T }{
	'*': {tt.Mul, tt.MulAssign},
	'/': {tt.Div, tt.DivAssign},
	'%': {tt.Mod, tt.ModAssign},
	'^': {tt.Xor, tt.XorAssign},
	'=': {tt.Assign, tt.Eq},
	'!': {tt.Not, tt.Neq},
}

var xeqOps = map[rune]*struct {
	t, eqt, xt tt.T
	x          rune
}{
	'+': {tt.Add, tt.AddAssign, tt.Inc, '+'},
	'-': {tt.Sub, tt.SubAssign, tt.Dec, '-'},
	'|': {tt.Or, tt.OrAssign, tt.Lor, '|'},
}

func (lx *Lexer) scanOperator(r rune) tt.T {
	s := lx.s

	if r == '\n' {
		lx.insertSemi = false
		return tt.Semi
	} else if r == '.' {
		if s.Scan('.') {
			if s.Scan('.') {
				return tt.Ellipsis
			}
			lx.reportf("two dots, expecting one more")
			return tt.Illegal
		}

		return tt.Dot
	} else if ret, found := simpleOps[r]; found {
		return ret
	} else if o, found := eqOps[r]; found {
		if s.Scan('=') {
			return o.eqt
		}
		return o.t
	} else if o, found := xeqOps[r]; found {
		if s.Scan(o.x) {
			return o.xt
		} else if s.Scan('=') {
			return o.eqt
		}
		return o.t
	}

	switch r {
	case '<':
		if s.Scan('=') {
			return tt.Leq
		} else if s.Scan('<') {
			if s.Scan('=') {
				return tt.ShiftLeftAssign
			}
			return tt.ShiftLeft
		}
		return tt.Less
	case '>':
		if s.Scan('=') {
			return tt.Geq
		} else if s.Scan('>') {
			if s.Scan('=') {
				return tt.ShiftRightAssign
			}
			return tt.ShiftRight
		}
		return tt.Greater
	case '&':
		if s.Scan('^') {
			if s.Scan('=') {
				return tt.NandAssign
			} else {
				return tt.Nand
			}
		}
		if s.Scan('=') {
			return tt.AndAssign
		} else if s.Scan('&') {
			return tt.Land
		}

		return tt.And
	}

	if !lx.illegal {
		lx.illegal = true
		lx.reportf("illegal character")
	}
	return tt.Illegal
}
