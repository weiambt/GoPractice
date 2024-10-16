package main

func main() {
	a := []int{1, 2, 3}
	sol(a)
	print(a[1]) //修改了，数组传参属于引用传参

	b := 1
	change(b)
	print(b) //输出原来值
}
func sol(arr []int) {
	arr[1] = 10
}

func change(a int) {
	a = 2
}
