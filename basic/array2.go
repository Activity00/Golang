package main

type Person struct {
	name string
}

func main() {
	p := Person{name: "wmh"}
	array0 := [...]Person{p}
	array1 := array0
	println(&array0, &array0[0])
	println(&array1, &array1[0])
}
