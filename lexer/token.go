package lexer

import (
	"fmt"

	"e8vm.net/leaf/lexer/token"
)

type Token struct {
	Token token.Token
	File  string
	Line  int
	Col   int
	Lit   string
}

func (t *Token) Clone() *Token {
	ret := new(Token)
	ret.File = t.File
	ret.Token = t.Token
	ret.Line = t.Line
	ret.Col = t.Col
	ret.Lit = t.Lit
	return ret
}

func (t *Token) Str(f string) string {
	if !t.Token.IsSymbol() {
		return fmt.Sprintf("%s:%d:%d: %s - %q",
			f, t.Line, t.Col,
			t.Token, t.Lit,
		)
	}

	return fmt.Sprintf("%s:%d:%d: %s", f, t.Line, t.Col, t.Token)
}

func (t *Token) String() string {
	if !t.Token.IsSymbol() {
		return fmt.Sprintf("%d:%d: %s - %q",
			t.Line, t.Col,
			t.Token, t.Lit,
		)
	}

	return fmt.Sprintf("%d:%d: %s", t.Line, t.Col, t.Token)
}
