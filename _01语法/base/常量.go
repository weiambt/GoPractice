package main

import "fmt"

func test_const() {
	//1.不指定类型
	const time = 1
	//2.指定类型
	const age int = 2

	//统一申明
	const (
		a        = 1
		b string = "a"
	)

}

func test_iota() {
	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
}

func test_iota2() {
	const (
		a = 1 << iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d)
}

func main() {
	test_const()
	test_iota()
}
