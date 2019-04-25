package main

import (
	"fmt"
)

func foo(slice []int) []int {
	return slice
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	slice1 := slice[1:3]
	println(len(slice1), cap(slice1))
	// 切片共享内存
	println(slice1[0])
	slice[1] = 999
	println(slice1[0])

	//追加 ...
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Print(s1)
	fmt.Print(s2)
	fmt.Print(append(s1, s2...))

	slice2 := make([]int, 1e6)
	slice3 := foo(slice2)
	fmt.Print(slice3)
}
