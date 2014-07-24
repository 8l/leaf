package tt

import (
	"fmt"
)

var strs = map[T]string{
	Illegal: "Illegal",
	Comment: "Comment",

	Ident:  "Ident",
	Int:    "Int",
	Float:  "Float",
	Char:   "Char",
	String: "String",

	Lparen: "(",
	Lbrace: "{",
	Lbrack: "[",
	Rparen: ")",
	Rbrace: "}",
	Rbrack: "]",

	Assign: "=",
	Dollar: "$",
	Colon:  ":",
	Semi:   ";",
	Comma:  ",",

	Const: "const",
	Var:   "var",
	Func:  "func",
}

func (t T) String() string {
	if s, found := strs[t]; found {
		return s
	}
	return fmt.Sprintf("<#%d>", int(t))
}

var reserves = func() map[string]T {
	ret := make(map[string]T)
	for i := keywordBegin + 1; i < keywordEnd; i++ {
		ret[strs[i]] = i
	}

	return ret
}()

// FromIdent returns the related keyword token if it is a keyword
// or type; returns Ident otherwise.
func FromIdent(s string) T {
	if i, found := reserves[s]; found {
		return i
	}
	return Ident
}
