package symbol

import (
	"e8vm.net/leaf/lexer"
)

type Symbol interface {
	Name() string        // a symbol must have a name
	Kind() Kind          // a symbol must also have a kind
	Token() *lexer.Token // a symbole also has a declared position
}
