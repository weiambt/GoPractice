package main

import (
	"encoding/json"
	"fmt"
)

func test1() {
	per := Per{id: 1, name: "a"}
	res, err := json.Marshal(per)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))
	ans := new(Per)
	json.Unmarshal(res, &ans)
	fmt.Println(ans)
}

func test2() {
	per := Per2{Id: 1, Name: "a"}
	res, err := json.Marshal(per)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(res))
	ans := new(Per2)
	json.Unmarshal(res, &ans)
	fmt.Println(ans)
}

func main() {
	//使用小写的字段就无法序列化
	test1()

	//使用大写的字段
	test2()
}

type Per struct {
	id   int
	name string
}

type Per2 struct {
	Id   int
	Name string
}
