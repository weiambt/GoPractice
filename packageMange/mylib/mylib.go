package mylib

import "fmt"

//先调用Init方法，再调用本身

func Solve() string {
	fmt.Println("solve")
	return "ok"
}

func init() {
	fmt.Println("init")
}
