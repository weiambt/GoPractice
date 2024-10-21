package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s := struct{}{}
	sz := unsafe.Sizeof(s)
	fmt.Println(sz)

	a := make([]byte, 10, 20)
	sz = unsafe.Sizeof(a)
	fmt.Println(sz)

	//int 在不同的平台上，占用的字节数不一样。
	// 32位平台上，int占用4个字节;64位平台上，int占用8个字节。
	var b int
	sz = unsafe.Sizeof(b)
	fmt.Println(sz)

	// int64、int32这些在不同的平台上，占用的字节数固定。
	var c int64
	sz = unsafe.Sizeof(c)
	fmt.Println(sz)
}
