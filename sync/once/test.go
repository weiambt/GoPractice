package main

import (
	"fmt"
	"sync"
)

//错误使用，Do中嵌套调用Do造成死锁

var msg string
var one sync.Once

func main() {
	one.Do(fun1)
}

func fun1() {
	fmt.Println("我是 fun1")
	one.Do(fun2)
}

func fun2() {
	fmt.Println("我是 fun2")
}
