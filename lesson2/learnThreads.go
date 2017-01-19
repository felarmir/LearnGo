package lesson2

import ("fmt"
  "time"
)

func ThreadLearn()  {
  go th("T1", 10)
  go th("T2", 10)
  go th("T3", 10)
  var input string
//===========================
  var c chan string = make(chan string)
  go pinger(c)
  go ponger(c)
  go printer(c)
//===========================
  fmt.Scanln(&input)
}

func th(name string, n int) {
  for i := 0;i < n;i++ {
    fmt.Println(name, "->", i)
  }
}

func pinger(c chan<- string) {
  for i := 0; ;i++ {
    c <- "ping"
  }
}

func ponger(c chan string) {
  for {
    c <- "pong"
    time.Sleep(time.Second * 2)
  }
}

func printer(c chan string)  {
  for {
    msg := <- c
    fmt.Println(msg)
    time.Sleep(time.Second * 1)
  }
}
