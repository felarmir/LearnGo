package lesson5

import (
	"flag"
	"fmt"
	"math/rand"
)

// go run commandLine.go -max=<value>
func main() {
	maxp := flag.Int("max", 6, "the max value")
	flag.Parse()
	fmt.Println(rand.Intn(*maxp))
}
