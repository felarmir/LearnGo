package designpatterns

import "sync"

type singleton struct {
  Name string
}

var instance *singleton
var once sync.Once

func GetInstanceSingletom() *singleton {
  once.Do(func() {
    instance = &singleton{}
  })
  return instance
}
