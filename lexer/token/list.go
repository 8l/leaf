package token

// List all the keywords
func Keywords() []Token { return keywordList }

// List all the operators
func Operators() []Token { return operatorList }

var keywordList = func() []Token {
	ret := make([]Token, 0, keywordEnd-keywordBegin-1)
	for i := keywordBegin + 1; i < keywordEnd; i++ {
		ret = append(ret, i)
	}
	return ret
}()

var operatorList = func() []Token {
	ret := make([]Token, 0, operatorEnd-operatorBegin-1)
	for i := operatorBegin + 1; i < operatorEnd; i++ {
		ret = append(ret, i)
	}
	return ret
}()
