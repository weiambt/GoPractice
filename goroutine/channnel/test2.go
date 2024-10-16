package main

import (
	"fmt"
	"time"
)

func test_wuhuanchong() {
	//无缓冲channel
	ch := make(chan int, 0)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println(len(ch), cap(ch)) //长度和容量
		}
	}()
	ans := <-ch
	fmt.Println(ans)
}

func main() {

	//test_wuhuanchong()

	test_range()
}

func test_range() {
	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	time.Sleep(time.Second)
	//如果不close channel就会报错
	for x := range ch {
		fmt.Println(x)
	}
}
