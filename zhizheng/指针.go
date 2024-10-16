package main

import "fmt"

func main() {
	a := 1
	var p *int
	p = &a
	fmt.Println(*p) //指针的值
	fmt.Println(p)
	fmt.Println(&a)

	//二级指针
	pp := &p
	fmt.Println(*pp)
	fmt.Println(pp)
	fmt.Println(&p)
}
