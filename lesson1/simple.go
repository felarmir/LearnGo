package lesson1

import "fmt"

func Simple1() {
	//values
	var m string
	var i = 10
	x := "string value"
	y := 10
	m = "Test 2"

	fmt.Println("M:", m, "|i:", i, "|y:", y, "|x:", x)
}

func Simple2() {
	//for and if
	fmt.Println("============ Simple 2 (for, if, switch) ====================")
	fmt.Println("---------------( for )-----------------")
	i := 1
	for i < 10 {
		fmt.Println("->", 1)
		i = i + 1
	}
	for i := 0; i < 10; i++ {
		fmt.Println("iter ", i)
	}
	fmt.Println("---------------( if )-----------------")
	x := 10
	if x > 5 {
		fmt.Println("Big number")
	} else {
		fmt.Println("Small number")
	}
	fmt.Println("---------------( switch )-----------------")
	switch x {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 10:
		fmt.Println("Ten")
	default:
		fmt.Println("Unknown Number")
	}
}

func Simple3() {
	// Arrays
	fmt.Println("============ Array ====================")
	fmt.Println("---------------( simple array init )-----------------")
	var x [5]int
	x[3] = 29
	fmt.Println("Add value by ID: ", x)

	fmt.Println("---------------( array element summ )-----------------")
	var j [5]float64
	j[0] = 10
	j[1] = 13
	j[2] = 23
	j[3] = 45
	j[4] = 32

	var totalSumm float64 = 0
	for i := 0; i < len(j); i++ {
		totalSumm += j[i]
	}
	fmt.Println("Summ:", totalSumm)

	fmt.Println("---------------( cut  )-----------------")
	//var x []float64
	xc := make([]float64, 5) // empty array
	fmt.Println("array cut:", xc)

	arr := []float64{1, 2, 3, 4, 5}
	xa := arr[2:5]
	fmt.Println("after cut:", xa)
	fmt.Println("---------------( append  )-----------------")
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5, 6, 7)
	fmt.Println("After append:", slice2)
}

func Simple4() {
	//Map
	fmt.Println("===============( MAP (key-string, value - int) )===============")
	x := make(map[string]int)
	x["key"] = 10
	x["home"] = 23
	fmt.Println("x[key]:", x["key"])
	fmt.Println("---------------( map (key-int, value-int)  )-----------------")
	y := make(map[int]int)
	y[1] = 100
	y[2] = 200
	delete(y, 2)
	fmt.Println("y[1]:", y[1])

	fmt.Println("---------------( map (key-string, value-string)  )-----------------")
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["Ne"] = "Neon"
	fmt.Println("element name:", elements["N"])

	elm := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
	}

	if el, ok := elm["H"]; ok {
		fmt.Println(el["name"], el["state"])
	}
}
