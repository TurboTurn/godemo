package main

import "fmt"

func main() {
	b := func(a, b int, z float64) bool {
		return a*b < int(z)
	}(3, 2, 7.1)
	fmt.Println(b)

	//Go语言中的闭包同样也会引用到函数外的变量。闭包的实现确保只要闭包还被使用，那么
	//被闭包引用的变量会一直存在 go学习笔记P60
	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()
}
