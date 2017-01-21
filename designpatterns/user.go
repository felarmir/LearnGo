package designpatterns

import "fmt"

func PrintSingletonName()  {
  u1 := GetInstanceSingletom()
  fmt.Println("Singleton Name:", u1.Name)
}
