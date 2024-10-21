package main

import "fmt"

func printLenCap(a []int) {
	fmt.Printf("len=%d, cap=%d\n %v\n", len(a), cap(a), a)
}

// 测试：子切片添加元素会导致底层被修改
func test_1() {

	nums := make([]int, 0, 8)
	nums = append(nums, 1, 2, 3, 4, 5)
	nums2 := nums[2:4]
	printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 5]
	printLenCap(nums2) // len: 2, cap: 6 [3 4]

	//append的时候会在当前切片的len后面添加元素
	nums2 = append(nums2, 50, 60)
	printLenCap(nums)  // len: 5, cap: 8 [1 2 3 4 50]
	printLenCap(nums2) // len: 4, cap: 6 [3 4 50 60]

}

// 陷阱：_切片、数组修改元素
func test_2() {
	var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	modify_arr(arr)
	fmt.Println(arr)

	my_slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	modify_slice(my_slice)
	fmt.Println(my_slice)
}

// append
func test_3() {

	//创建切片时会有初始元素，append会继续添加在len的后面
	sli := make([]int, 2, 6) //[0 0]
	sli = append(sli, 1, 2)  // [0 0 1 2]
	printLenCap(sli)

	sli = append(sli, 3, 4) //[0 0 1 2 3 4]
	printLenCap(sli)

}

func modify_arr(a [10]int) {
	fmt.Println(a)
	a[0] = 100
}

func modify_slice(a []int) {
	a[0] = 100
}

func main() {
	//test_1()

	test_2()

	//test_3()

	//test_4()

	//test_5()
}
