// Package tok defines a token structure.
package tok

import (
	"fmt"
)

// Type is a token type.
type Type interface {
	String() string
	Code() int
	IsSymbol() bool // A symbol lacks a meaningful literal
}

// Token is a structure that describes a token in a file.
type Token struct {
	Type Type
	File string
	Line int
	Col  int
	Lit  string
}

// Clone clones the token and returns a copy.
func (t *Token) Clone() *Token {
	ret := new(Token)
	ret.Type = t.Type
	ret.File = t.File
	ret.Line = t.Line
	ret.Col = t.Col
	ret.Lit = t.Lit
	return ret
}

func (t *Token) String() string {
	var ret string
	ret = fmt.Sprintf("%d:%d %s", t.Line, t.Col, t.Type.String())
	if t.File != "" {
		ret = t.File + ":" + ret
	}

	if !t.Type.IsSymbol() {
		ret += fmt.Sprintf(" - %q", t.Lit)
	}
	return ret
}

// Is tests if a token is a type of token.
func (t *Token) Is(typ Type) bool {
	return t.Type.Code() == typ.Code()
}
