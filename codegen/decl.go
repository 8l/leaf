package codegen

import (
	"e8vm.net/leaf/ir/symbol"
	"e8vm.net/leaf/lexer"
)

type decl struct {
	class symbol.Class
	name  string
	pos   *lexer.Token
}
