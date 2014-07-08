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

	Add: "+",
	Sub: "-",
	Mul: "*",
	Div: "/",
	Mod: "%",

	And:        "&",
	Or:         "|",
	Xor:        "^",
	ShiftLeft:  "<<",
	ShiftRight: ">>",
	Nand:       "&^",

	AddAssign: "+=",
	SubAssign: "-=",
	MulAssign: "*=",
	DivAssign: "/=",
	ModAssign: "%=",

	AndAssign:        "&=",
	OrAssign:         "|=",
	XorAssign:        "^=",
	ShiftLeftAssign:  "<<=",
	ShiftRightAssign: ">>=",
	NandAssign:       "&^=",

	Land: "&&",
	Lor:  "||",
	Inc:  "++",
	Dec:  "--",

	Eq:      "==",
	Less:    "<",
	Greater: ">",
	Assign:  "=",
	Not:     "!",

	Neq:      "!=",
	Leq:      "<=",
	Geq:      ">=",
	Ellipsis: "...",

	Lparen: "(",
	Lbrack: "[",
	Lbrace: "{",
	Comma:  ",",
	Dot:    ".",

	Rparen: ")",
	Rbrack: "]",
	Rbrace: "}",
	Semi:   ";",
	Colon:  ":",

	Break:       "break",
	Case:        "case",
	Const:       "const",
	Continue:    "continue",
	Default:     "default",
	Else:        "else",
	Fallthrough: "fallthrough",
	For:         "for",
	Func:        "func",
	Goto:        "goto",
	If:          "if",
	Import:      "import",
	Return:      "return",
	Struct:      "struct",
	Switch:      "switch",
	Type:        "type",
	Var:         "var",
}

func (t T) String() string {
	if s, found := strs[t]; found {
		return s
	}
	return fmt.Sprintf("<na-%d>", int(t))
}

var keywords = func() map[string]T {
	ret := make(map[string]T)
	for i := keywordBegin + 1; i < keywordEnd; i++ {
		s := strs[i]
		ret[s] = i
	}
	return ret
}()

// Returns the related keyword token if it is a keyword; returns Ident otherwise.
func FromIdent(s string) T {
	if i, found := keywords[s]; found {
		return i
	}
	return Ident
}
