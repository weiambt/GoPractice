package main

import "fmt"

func main() {
	//初始化
	//mp := make(map[string]int)

	//指定长度
	//mp2 := make(map[string]int, 2)

	//mp3 := map[string]int{"a": 1, "b": 2}

	//delete(mp3, "a")

	//fmt.Println(mp3)

	mp := make(map[string]Per)

	mp["a"] = Per{Name: "zhangsan"}

	//编译报错
	//mp["a"].Name = "lisi"

	tmp := mp["a"]
	tmp.Name = "lisi"
	fmt.Println(mp) //zhangsan
	mp["a"] = tmp   //整体赋值
	fmt.Println(mp) //lisi
}

type Per struct {
	Name string
}
