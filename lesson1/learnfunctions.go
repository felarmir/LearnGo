package lesson1

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func twoValues(x int, y int) (int, int) {
	x = x * 3
	y = y * 3
	return x, y
}

func add(arg ...int) int {
	sum := 0
	for _, v := range arg {
		sum += v
	}
	return sum
}

func funcInFunc() {
	addF := func(x int, y int) int {
		return x + y
	}
	fmt.Println("fun In func:", addF(2, 2))
}

func TestFunc() {
	fmt.Println("===============( learn func )================")
	xs := []float64{23, 43, 12, 76, 3, 90, 43}
	fmt.Println("Average:", average(xs))
	fmt.Println("===============( learn func return two values )================")
	x, y := twoValues(3, 5)
	fmt.Printf("X:%d Y:%d \n", x, y)
	fmt.Println("===============( learn func add many values )================")
	fmt.Println("->", add(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println("===============( learn func in func )================")
	funcInFunc()
}
