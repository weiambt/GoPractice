package main

import "fmt"

func main() {
	mp := make(map[string]int)
	mp["1"] = 1
	mp["2"] = 2
	fmt.Println(mp)

	mp2 := make(map[string]map[string]int)
	mp2["1"] = map[string]int{"a": 1, "b": 2}
	fmt.Println(mp2) //map[1:map[a:1 b:2]]

}
