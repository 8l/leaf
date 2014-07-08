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
	Rparen: ")",
	Rbrace: "}",

	Assign: "=",
	Dollar: "$",
	Colon:  ":",
	Semi:   ";",
	Comma:  ",",

	Const: "const",
	Var:   "var",
	Func:  "func",

	U8:  "u8",
	I8:  "i8",
	U16: "u16",
	I16: "i16",
	U32: "u32",
	I32: "i32",
	F64: "f64",
	Str: "str",
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

	for i := typeBegin + 1; i < typeEnd; i++ {
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
