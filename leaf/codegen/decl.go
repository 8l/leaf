package codegen

import (
	"e8vm.net/leaf/leaf/ir/symbol"
	"e8vm.net/leaf/tools/tok"
)

type decl struct {
	class symbol.Class
	name  string
	pos   *tok.Token
}
