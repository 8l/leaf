// Package tt defines the tokens for the E8 assembly language
package tt

// T is the E8 assembly token type
type T int

// All tokens used in an E8 assembly file
const (
	// misc

	Illegal = T(iota + 1)
	Comment

	// literals

	literalBegin
	Ident
	Int
	Float
	Char
	String
	literalEnd

	// operators

	operatorBegin

	Lparen // (
	Lbrace // {
	Rparen // )
	Rbrace // }

	Assign // =
	Dollar // $
	Colon  // :
	Semi   // ;
	Comma  // ,

	operatorEnd

	// keywords

	keywordBegin
	Const
	Var
	Func
	keywordEnd

	// types

	typeBegin
	U8
	I8
	U16
	I16
	U32
	I32
	F64
	Str
	typeEnd
)

// Code returns the integer code for the token type.
func (t T) Code() int { return int(t) }

// IsOperator tests if it is an operator token type.
func (t T) IsOperator() bool {
	return operatorBegin < t && t < operatorEnd
}

// IsKeyword tests if it is a keyword token type.
func (t T) IsKeyword() bool {
	return keywordBegin < t && t < keywordEnd
}

// IsType tests if it is a type token type.
func (t T) IsType() bool {
	return typeBegin < t && t < typeEnd
}

// IsReserved tests if it is a reserved ident (a keyword or a type).
func (t T) IsReserved() bool {
	return t.IsKeyword() || t.IsType()
}

// IsLiteral tests if it is a literal type.
func (t T) IsLiteral() bool {
	return literalBegin < t && t < literalEnd
}

// IsSymbol tests if it is a symbol where its literal can be predicted.
func (t T) IsSymbol() bool {
	if t.IsLiteral() {
		return false
	}
	if t == Comment {
		return false
	}
	if t == Illegal {
		return false
	}
	return true
}
