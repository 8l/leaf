package build

import (
	"e8vm.net/leaf/tools/tok"
)

type label struct {
	pos *tok.Token // the position of the label
	loc int        // the position of the label
}
