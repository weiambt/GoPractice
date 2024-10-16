package main

func main() {
	//这是数组
	//a := [...]int{1, 2}
	//print(a)

	//方法1
	var a []int
	a = []int{1, 2, 3}

	//方法2
	var a []int
	a = make([]int, 5) //不指定cap
	a = make([]int, 5, 10)

	//方法3
	var a []int = make([]int, 5, 10)

	//方法4
	a := make([]int, 5, 10)

	print(a[0])
}
