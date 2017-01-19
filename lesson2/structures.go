package lesson2

import "fmt"

type Rectangle struct {
  A float64
  B float64
}

func (r *Rectangle) GetArea () float64  {
  return r.A * r.B
}

type Person struct {
  Name string
}

func (p *Person) SayName() {
  fmt.Println("My name is ", p.Name)
}

func (p *Person) SetName(name string) {
  p.Name = name
}

type User struct {
  Self Person
  Email string
}

//==================================
type Shape interface {
  GetArea() float64
}

func TotalArea(shape ...Shape) float64 {
  var area float64
  for _, s := range shape {
    area += s.GetArea()
  }
  return area
}

func TestInterface() {
  r1 := Rectangle{23, 32}
  r2 := Rectangle{12, 32}
  r3 := Rectangle{1, 3}
  r4 := Rectangle{3, 54}
  r5 := Rectangle{2, 31}
  fmt.Println("Summ area:", TotalArea(&r1, &r2, &r3, &r4, &r5))
}
