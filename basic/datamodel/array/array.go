package main

func printArray(array *[10]int, count int) {
	for i := 0; i < count; i++ {
		println(array[i])
	}
}

func main() {
	// 申明与初始化
	var array0 [10]int
	for i, x := range array0 {
		println(i, x)
	}

	// 初始化
	array1 := [10]int{1, 2, 3, 4, 5}
	printArray(&array1, len(array1))

	// 默认长度
	array2 := [...]int{1, 2, 4}
	for i, x := range array2 {
		println(i, x)
	}

	array3 := [...]int{1, 2, 4}
	for i, x := range array3 {
		println(i, x)
	}

	//特定赋值
	array4 := [...]int{1: 100, 5: 500}
	for i, x := range array4 {
		println(i, x)
	}

	println(array4[1])
	array4[1] = 1000
	println(array4[1])

	//指针数组
	array5 := [5]*int{1: new(int), 4: new(int)}
	println(array5[1], array5[3])
}
