package main

type Person struct {
	name string
}

func main() {
	// 都会值拷贝
	p := Person{name: "wmh"}
	array0 := [1]Person{p}
	var array1 [1]Person
	array1 = array0

	println(&array0[0])
	println(&array1[0])
}
