package lesson2

import "testing"

type testPair struct {
  r Rectangle
  result float64
}

var tests = []testPair {
  {Rectangle{2, 2}, 4},
  {Rectangle{3, 3}, 9},
  {Rectangle{22, 10}, 220},
  {Rectangle{8, 8}, 64},
}

func TestAverage(t *testing.T) {
  for _, s := range tests {
    v := TotalArea(&s.r)
    if v != s.result {
      t.Error("Except ", s.result, ", got", v)
    }
  }
}
