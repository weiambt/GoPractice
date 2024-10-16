package main

func change(arr [2]int) {
	arr[0] = 100
}

func change2(arr *[2]int) {
	arr[0] = 200
}

func main() {
	a := [...]int{1, 2}
	change2(&a) //[200 2]
}
