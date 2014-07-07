package tt

// List all the keywords
func Keywords() []T { return keywordList }

// List all the operators
func Operators() []T { return operatorList }

var keywordList = func() []T {
	ret := make([]T, 0, keywordEnd-keywordBegin-1)
	for i := keywordBegin + 1; i < keywordEnd; i++ {
		ret = append(ret, i)
	}
	return ret
}()

var operatorList = func() []T {
	ret := make([]T, 0, operatorEnd-operatorBegin-1)
	for i := operatorBegin + 1; i < operatorEnd; i++ {
		ret = append(ret, i)
	}
	return ret
}()
