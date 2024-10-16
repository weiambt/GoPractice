package main

import "fmt"

func test1() {
	s := make([]int, 0, 10)
	sub := []int{1, 2}
	s = append(s, sub...) //将sub打散成一个个元素，并且append也就是接收可变参数

}

func test2() {
	s := make([]int, 0, 10)
	//可变参数，自定义实现append
	myAppend := func(s []int, args ...int) []int {
		s = append(s, args...)
		return s //如果s不加指针就要返回
	}
	myAppend(s, -1, -2)
	fmt.Println(s)

}
func main() {
	test1()
	test2()
}
