package tt

func (t T) IsOperator() bool {
	return operatorBegin < t && t < operatorEnd
}

func (t T) IsKeyword() bool {
	return keywordBegin < t && t < keywordEnd
}

func (t T) IsLiteral() bool {
	return literalBegin < t && t < literalEnd
}

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
