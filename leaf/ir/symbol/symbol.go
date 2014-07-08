package symbol

type Symbol struct {
	name   string
	class  Class
	object interface{}
}

func (self *Symbol) Class() Class     { return self.class }
func (self *Symbol) Name() string     { return self.name }
func (self *Symbol) Obj() interface{} { return self.object }
