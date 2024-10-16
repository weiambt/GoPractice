package main

import "fmt"

func main() {
	mp := make(map[string][]int, 1)
	mp["a"] = []int{1, 2, 3}
	var ans []int
	ans = append(ans, mp["a"]...)

	fmt.Println(ans)
}
