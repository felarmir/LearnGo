package main

import "./designpatterns"

//"fmt"

func main() {
	s := designpatterns.GetInstanceSingletom()
	var name = "Denis"
	s.Name = name
	designpatterns.PrintSingletonName()
}
