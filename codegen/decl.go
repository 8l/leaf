package codegen

import (
	"e8vm.net/leaf/ir/symbol"
	"e8vm.net/util/tok"
)

type decl struct {
	class symbol.Class
	name  string
	pos   *tok.Token
}
