package designpatterns

type Prototyper interface{
  Clone() Prototyper
  GetName() string
}

type ConcretProductA struct {
  name string
}

func (self *ConcretProductA) GetName() string {
  return self.name
}

func (self *ConcretProductA) Clone() Prototyper {
  return &ConcretProductA{self.name}
}

type ConcretProductB struct {
  name string
}

func (self *ConcretProductB) GetName() string {
  return self.name
}

func (self *ConcretProductB) Clone() Prototyper {
  return &ConcretProductB{self.name}
}
