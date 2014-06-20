package lexer

import (
	t "e8vm.net/leaf/lexer/token"
)

var simpleOps = map[rune]t.Token{
	':': t.Colon,
	',': t.Comma,
	';': t.Semi,
	'(': t.Lparen,
	')': t.Rparen,
	'[': t.Lbrack,
	']': t.Rbrack,
	'{': t.Lbrace,
	'}': t.Rbrace,
}

var eqOps = map[rune]*struct{ t, eqt t.Token }{
	'*': {t.Mul, t.MulAssign},
	'/': {t.Div, t.DivAssign},
	'%': {t.Mod, t.ModAssign},
	'^': {t.Xor, t.XorAssign},
	'=': {t.Assign, t.Eq},
	'!': {t.Not, t.Neq},
}

var xeqOps = map[rune]*struct {
	t, eqt, xt t.Token
	x          rune
}{
	'+': {t.Add, t.AddAssign, t.Inc, '+'},
	'-': {t.Sub, t.SubAssign, t.Dec, '-'},
	'|': {t.Or, t.OrAssign, t.Lor, '|'},
}

func (self *Lexer) scanOperator(r rune) t.Token {
	s := self.s

	if r == '\n' {
		self.insertSemi = false
		return t.Semi
	} else if r == '.' {
		if s.Scan('.') {
			if s.Scan('.') {
				return t.Ellipsis
			}
			self.failf("two dots, expecting one more")
			return t.Illegal
		}

		return t.Dot
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
			return t.Leq
		} else if s.Scan('<') {
			if s.Scan('=') {
				return t.ShiftLeftAssign
			}
			return t.ShiftLeft
		}
		return t.Less
	case '>':
		if s.Scan('=') {
			return t.Geq
		} else if s.Scan('>') {
			if s.Scan('=') {
				return t.ShiftRightAssign
			}
			return t.ShiftRight
		}
		return t.Greater
	case '&':
		if s.Scan('^') {
			if s.Scan('=') {
				return t.NandAssign
			} else {
				return t.Nand
			}
		}
		if s.Scan('=') {
			return t.AndAssign
		} else if s.Scan('&') {
			return t.Land
		}

		return t.And
	}

	if !self.illegal {
		self.illegal = true
		self.failf("illegal character")
	}
	return t.Illegal
}
