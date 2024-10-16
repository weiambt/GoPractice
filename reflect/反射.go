package main

import (
	"fmt"
	"reflect"
)

type A struct {
	Name string
	Age  int
}

func (a A) check() {
	fmt.Println("check")
}

func solve(arg interface{}) {
	//获取对象类型
	inType := reflect.TypeOf(arg)
	//获取对象值
	inVal := reflect.ValueOf(arg)
	fmt.Println(inType)
	fmt.Println(inVal)

	//获取所有变量类型
	for i := 0; i < inType.NumField(); i++ {
		inTypeField := inType.Field(i)
		inValField := inVal.Field(i)
		fmt.Println("---", inTypeField.Name, inValField)
		//区别
		fmt.Println(inVal.Field(i))
		fmt.Println(inVal.Field(i).Interface())
	}
	fmt.Println(inType.NumMethod())

	//获取所有方法
	//todo:   不知道为什么不输出？？？？？？
	for i := 0; i < inType.NumMethod(); i++ {
		inTypeMethod := inType.Method(i)
		fmt.Println("---", inTypeMethod.Name, inTypeMethod)
	}

}

func main() {
	ans := A{Name: "zjhang", Age: 18}
	solve(ans)
}
