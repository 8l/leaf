package tt

// List all the keywords
func Keywords() []T { return keywordList }

// List all the operators
func Operators() []T { return operatorList }

func makeList(from, to T) []T {
	ret := make([]T, 0, to-from-1)
	for i := from + 1; i < to; i++ {
		ret = append(ret, i)
	}
	return ret
}

var (
	keywordList  = makeList(keywordBegin, keywordEnd)
	operatorList = makeList(operatorBegin, operatorEnd)
)
