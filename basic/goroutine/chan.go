package main

func main() {
	unbuffered := make(chan int)
	buffered := make(chan string, 10) // 10 个值的缓冲
}
