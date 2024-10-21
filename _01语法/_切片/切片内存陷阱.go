package main

import "fmt"

// append 超过cap的容量，会重新分配内存
func test_4() {
	sli := make([]int, 0, 2)
	sli = append(sli, 1, 2)
	sub := sli[1:]
	fmt.Println(sub)

	//超过cap，创建了新的底层数组，之前的底层数组被sub指向
	sli = append(sli, 3, 4)
	//修改原切片，sli不会被影响，只会影响sub
	sub[0] = 100
	fmt.Println(sli) // [1 2 3 4]
	fmt.Println(sub) // [100]
}

func test_5() {
	//不超过cap,超过len，还是同一个数组
	sli := make([]int, 0, 4)
	sli = append(sli, 1, 2)
	sub := sli[1:]

	sli = append(sli, 3)
	sli[1] = 100
	fmt.Println(sli) // [1 100 3]
	fmt.Println(sub) // [100]
}

func main() {
	test_4()
	test_5()
}
