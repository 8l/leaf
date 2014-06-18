package comperr

import (
	"fmt"
)

type Error struct {
	Err  error
	File string
	Line int
	Col  int
}

func (e *Error) Error() string {
	prefix := ""
	if e.File != "" {
		prefix = e.File + ":"
	}
	return fmt.Sprintf("%s%d:%d: %s",
		prefix, e.Line, e.Col, e.Err,
	)
}
