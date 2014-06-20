package token

func (t Token) IsOperator() bool {
	return operatorBegin < t && t < operatorEnd
}

func (t Token) IsKeyword() bool {
	return keywordBegin < t && t < keywordEnd
}

func (t Token) IsLiteral() bool {
	return literalBegin < t && t < literalEnd
}

func (t Token) IsSymbol() bool {
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
