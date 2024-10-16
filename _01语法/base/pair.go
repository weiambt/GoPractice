package main

import "fmt"

func solve() (int, int) {
	return 1, 0
}

func main() {
	var a interface{}
	a = 1
	//方式1
	b, err := a.(int)
	if err {
		fmt.Println(err)
	}
	fmt.Println(b)

	//方式2
	c := a.(string)
	fmt.Printf("%T\n", c)
}
