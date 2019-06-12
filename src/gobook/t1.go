package main

import "fmt"

func main() {
	f := func(x, y int) int { //匿名函数，可以随意对该匿名函数变量进行传递和调用
		return x + y
	}
	fmt.Println(f(4, 5))

	var ifly iFly = new(Bird)
	ifly.Fly()
}

type Bird struct {
}

func (b Bird) Fly() {
	fmt.Println("鸟飞行")
}

type iFly interface {
	Fly()
}
