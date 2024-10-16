package main

import "fmt"

// 计算距离
func (s *Point) Dist(arg Point) float64 {
	return arg.Y - s.Y + arg.X - s.X
}

func main() {
	a := Point{X: 1, Y: 1}
	b := Point{X: 0, Y: 0}

	//定义方法值
	method_a := a.Dist
	fmt.Println(method_a(b))
}

type Point struct {
	X, Y float64
}
