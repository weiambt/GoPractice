package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 100
	}()
	ans := <-ch
	fmt.Println(ans)
}
