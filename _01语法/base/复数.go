package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a complex64
	a = 1 + 10i
	fmt.Println(a)

	b := 2 + 10i + a
	fmt.Println(b)

	strconv.Itoa(23)

}
