package exprs

import (
	"e8vm.net/leaf/codegen/types"
)

// expression result
// often a location in memory
// a stack location, or an absolute location
type Expr interface{}

type IntExpr struct {
	Type  types.Type
	Value int64 // should be good for int values
}
