package designpatterns

import "log"

type Creater interface {
  CreateProduct(act string) Producer
  registerProduct(Producer)
}

type Producer interface {
  Use() string
}

type ConcrateCreat struct {
  proucts []*Producer
}

func (self *ConcrateCreat) registerProduct(product Producer) {
  self.proucts = append(self.proucts, &product)
}

type ConcrateProductA struct {
  act string
}

func (self *ConcrateProductA) Use() string {
  return self.act
}

type ConcrateProductB struct {
  act string
}

func (self *ConcrateProductB) Use() string {
  return self.act
}

type ConcrateProductC struct {
  act string
}

func (self *ConcrateProductC) Use() string {
  return self.act
}

func (self *ConcrateCreat) CreateProduct(act string) Producer {
  var product Producer
  switch act {
  case "A":
    product = &ConcrateProductA{act}
  case "B":
    product = &ConcrateProductB{act}
  case "C":
    product = &ConcrateProductC{act}
  default:
    log.Fatalln("Unknown Error")
  }
  self.registerProduct(product)
  return product
}
