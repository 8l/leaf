package types

import (
	"fmt"
)

type Named struct {
	Path string
	Name string
	Type Type
}

func (self *Named) Size() uint32 {
	return self.Type.Size()
}

func (self *Named) String() string {
	return fmt.Sprintf("%q.%s", self.Path, self.Name)
}

