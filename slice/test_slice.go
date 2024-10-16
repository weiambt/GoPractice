package main

import (
	"fmt"
	"time"
)

func sol1(N int, size int) {
	start := time.Now()
	for n := 0; n < N; n++ {
		data := make([]int, 0, size) //指定容量
		for k := 0; k < size; k++ {
			data = append(data, k)
		}
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func sol2(N int, size int) {
	start := time.Now()
	for n := 0; n < N; n++ {
		data := make([]int, 0) //不指定
		for k := 0; k < size; k++ {
			data = append(data, k)
		}
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func main() {
	//测试1
	N := int(1e7)
	size := 1000
	//sol1(N, size)
	//sol2(N, size)
	//9.901573417s
	//22.209529542s

	//2
	N = int(1e4)
	size = 10000
	sol1(N, size)
	sol2(N, size)

}
