package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// 必需要定义一个局部变量避免闭包重复利用相同的i
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
