package main

import (
	"bytes"
	"fmt"
	"strings"
)

func test1() {
	var s strings.Builder
	s.WriteString("helloworld")
	s.WriteString("as")
	fmt.Println(s.String())
}

func test2() {
	var s bytes.Buffer
	s.WriteString("h")
	s.WriteString("hs")
	fmt.Println(s.String())
}

func test3() {
	s := make([]byte, 0, 10)
	a := "1"
	//s = append(s, a) //报错，要写...
	s = append(s, a...) //

	fmt.Println(string(s))
}
func byteConcat(n int, str string) string {
	buf := make([]byte, 0)

	buf = append(buf, str...)

	return string(buf)
}

func main() {
	//test1()
	test2()
}
