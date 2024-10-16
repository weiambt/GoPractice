package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	a int
	b string
	c float64
}

func main() {
	data := `{
		 "a":1,
		 "b":<,
		 "c":15.15
	 }`
	ans := &A{}
	err := json.Unmarshal([]byte(data), &ans)
	fmt.Println(err)
	fmt.Println(ans)

	//js[0] = 11
	//json.Unmarshal(js, &ans)
	//fmt.Println(ans)

	//json.Unmarshal(data, &ans)

}
