package designpatterns

import "testing"

func TestPrototype(t *testing.T) {
  productA := ConcretProductA{"A"}
  productB := ConcretProductB{"B"}

  cloneProductA := productA.Clone()
  cloneProductB := productB.Clone()

  if cloneProductA.GetName() != productA.GetName() {
    t.Error("Error \"A\" is not equal")
  }
  if cloneProductB.GetName() != productB.GetName() {
    t.Error("Error \"B\" is not equal")
  }

}
