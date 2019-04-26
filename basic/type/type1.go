package main

type IP []byte

func test(p []int) {
	p[0] = 999
}

func main() {
	ip := IP{}
	println(ip)

	p := []int{1}
	test(p)
	println(p[0])
}
