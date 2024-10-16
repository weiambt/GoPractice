package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
			fmt.Println("----", len(ch))
		}
	}()

	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}
