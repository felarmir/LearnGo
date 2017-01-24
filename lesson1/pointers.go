package lesson1

import "fmt"

func zero(v *int) {
	*v = 0
}

func one(v *int) {
	*v = 1
}

func PointerTest2() {
	v := new(int)
	one(v)
	fmt.Println("Pointer test2:", *v)
}

func PointerTest() {
	x := 65
	zero(&x)
	fmt.Println("Pointer test:", x)
}

func setMap(data *interface{}) {
	mp := make(map[string]string)
	mp["name"] = "Denis"
	mp["job"] = "Developer"
	*data = mp
}

func PointerInterfaceTest() {
	var data interface{}
	setMap(&data)
	d := data.(map[string]string)
	fmt.Println(d["name"])
}