package codegen

import (
	"strconv"
)

func parseInt(s string) (int64, error) {
	return strconv.ParseInt(s, 0, 32)
}

func parseStr(s string) (string, error) {
	return strconv.Unquote(s)
}
