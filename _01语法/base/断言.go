package main

import "fmt"

func check(arg interface{}) {
	//断言arg的类型是string，如果不是就报错，如果是就转成string
	a := arg.(string)
	fmt.Printf("%T\n", a)

	fmt.Println(a)
}

func main() {
	//check(1)    //panic: interface conversion: interface {} is int, not string

	check("ac")
	//string
	//ac
}
