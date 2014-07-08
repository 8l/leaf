package tt

// Keywords lists all the keyword types
func Keywords() []T { return keywordList }

// Operators lists all the operator types
func Operators() []T { return operatorList }

// Types lists all the type types
func Types() []T { return typeList }

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
	typeList     = makeList(typeBegin, typeEnd)
)
