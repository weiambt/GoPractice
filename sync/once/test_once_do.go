package main

import (
	"fmt"
	"sync"
	"time"
)

//sync.Once.Do方法
//只执行一次

var msg2 string

func main() {
	var one sync.Once
	for i := 0; i < 5; i++ {
		go func(i int) {
			one.Do(func() {
				fmt.Printf("%d 执行初始化！\n", i)
				msg2 = "Your Need Data"
			})
			fmt.Println(msg2)
		}(i)
	}
	time.Sleep(3 * time.Second)
}
