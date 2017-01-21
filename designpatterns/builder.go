package designpatterns

type Builder interface {
  makeHeader(str string)
  makeContent(str string)
  makeFooter(str string)
}

type Product struct {
  Header string
  Content string
  Footer string
}

type Director struct {
  builder Builder
}

func (self *Director) Construct() {
  self.builder.makeHeader("Header")
  self.builder.makeContent("Content")
  self.builder.makeFooter("Footer")
}

type PersonBuilder struct {
  product *Product
}

func (self *PersonBuilder) makeHeader(str string) {
  self.product.Header = "<header>" + str + "</header>\n"
}

func (self *PersonBuilder) makeContent(str string) {
  self.product.Content = "<content>" + str + "</content>\n"
}

func (self *PersonBuilder) makeFooter(str string) {
  self.product.Footer = "<footer>" + str + "</footer>\n"
}

func (self *Product) Show() string {
  var result string
  result = self.Header + self.Content + self.Footer
  return result
}
