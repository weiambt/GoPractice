package main

import "fmt"

func main() {
	ch := make(chan int, 12)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for {
		//ok为true说明channel没有关闭，为false说明管道已经关闭
		if val, ok := <-ch; ok {
			fmt.Println(val)
		} else {
			break
		}
	}

}
