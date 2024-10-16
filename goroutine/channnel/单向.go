package main

import "fmt"

func input(in chan<- int) {
	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)
}

func output(out <-chan int) {
	for x := range out {
		fmt.Println(x)
	}
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-out)
	//}
}

func main() {
	ch := make(chan int, 10)

	go input(ch)
	output(ch)
}
