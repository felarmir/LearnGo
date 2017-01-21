package designpatterns

import "testing"

func TestBuilder(t *testing.T) {
  res := "<header>Header</header>\n"+
          "<content>Content</content>\n"+
          "<footer>Footer</footer>\n"

  product := &Product{}
  director := Director{&PersonBuilder{product}}
  director.Construct()
  result := product.Show()
  if result != res {
    t.Errorf("Except result to %s, but %s", result, res)
  }
}
