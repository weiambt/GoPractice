package main

import (
	"fmt"
	"strconv"
)

func solve(a interface{}) {
	b := a.(string)
	fmt.Println(b)

	c := strconv.Itoa(3)
	fmt.Println(c)

	d, err := strconv.Atoi("32")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func main() {
	solve("ada")
}
