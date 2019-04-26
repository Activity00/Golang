package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter   int
	waitGroup sync.WaitGroup
)

func main() {
	waitGroup.Add(2)

	go incCounter(1)
	go incCounter(2)

	waitGroup.Wait()
	fmt.Println("counter", counter)
}

func incCounter(id int) {
	defer waitGroup.Done()
	for count := 0; count < 2; count++ {
		value := counter
		// 当前goroutine 从线程退出，并且放到队列
		runtime.Gosched()
		value++
		counter = value
	}
}
