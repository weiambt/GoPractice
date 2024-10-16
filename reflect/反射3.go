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
	// 获取对象类型
	inType := reflect.TypeOf(arg)
	// 获取对象值
	inVal := reflect.ValueOf(arg)
	fmt.Println("类型:", inType)
	fmt.Println("值:", inVal)

	// 获取所有变量类型
	for i := 0; i < inType.NumField(); i++ {
		inTypeField := inType.Field(i)
		inValField := inVal.Field(i)
		fmt.Println("--- 字段名:", inTypeField.Name, "值:", inValField)
		fmt.Println("字段值:", inVal.Field(i))
		fmt.Println("字段接口值:", inVal.Field(i).Interface())
	}

	fmt.Println("方法数量:", inType.NumMethod())
	// 获取所有方法
	for i := 0; i < inType.NumMethod(); i++ {
		inTypeMethod := inType.Method(i)
		fmt.Println("--- 方法名:", inTypeMethod.Name, "类型:", inTypeMethod.Type)
	}
}

func main() {
	ans := A{Name: "zjhang", Age: 18}
	solve(ans)
}
