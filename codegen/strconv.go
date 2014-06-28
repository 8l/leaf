package codegen

import (
	"fmt"
	"math"
	"strconv"
)

func unquoteChar(s string) (uint8, error) {
	ef := fmt.Errorf

	n := len(s)
	if n < 3 {
		return 0, ef("invalid char literal")
	}
	if s[0] != '\'' || s[n-1] != '\'' {
		return 0, ef("invalid quoting char literal")
	}

	s = s[1 : n-1]
	ret, multi, tail, err := strconv.UnquoteChar(s, '\'')
	if multi {
		return 0, ef("multibyte char not allowed")
	} else if tail != "" {
		return 0, ef("char lit has a tail")
	} else if err != nil {
		return 0, ef("invalid char literal: %s, %v", s, err)
	} else if ret > math.MaxUint8 || ret < 0 {
		return 0, ef("invalid char value")
	}

	return uint8(ret), nil
}
