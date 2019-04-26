package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//分配一个逻辑处理器给调读者使用
	runtime.GOMAXPROCS(1)
	//计数+2 等待来两个groutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start Goroutines")
	go func() {
		//函数退出时调用Done来通知main函数工作完成
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		//函数退出时调用Done来通知main函数工作完成
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("Waiting for finish")
	wg.Wait()
	fmt.Println("finished")

}
