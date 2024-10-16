package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		//如果没有default语句，那么select语句将被阻塞，直到至少有一个通信可以进行下去
		//并且如果没有default语句，那么select语句将被阻塞，直到至少有一个通信可以进行下去。
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
			time.Sleep(time.Second)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
