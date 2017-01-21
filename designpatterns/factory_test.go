package designpatterns

import "testing"

func TestFactory(t *testing.T) {
  actions := []string{"A", "B", "C"}
  factory := new(ConcrateCreat)
  products := []Producer {
    factory.CreateProduct("A"),
    factory.CreateProduct("B"),
    factory.CreateProduct("C"),
  }
  for i, product := range products {
    if act := product.Use(); act != actions[i] {
      t.Errorf("Expect action to %s, but %s.\n", act[i], act)
    }
  }
}
