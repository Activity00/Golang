package main

func printSlice(slice []string) {
	for _, x := range slice {
		println(x)
	}
}

func main() {
	// 创建切片0
	slice := make([]int, 3, 5) // 长度和容量
	for _, x := range slice {
		println(x)
	}

	// 创建切片1
	slice1 := []string{"red", "blue", "green", "yellow", "pink"}
	printSlice(slice1)

	slice2 := []string{99: ""}
	printSlice(slice2)
	// nil 切片
	var slice3 []string
	printSlice(slice3)
	println(slice3 == nil)

	println(len(slice), cap(slice))

}
