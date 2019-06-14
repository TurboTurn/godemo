package main

import (
	"fmt"
)

//当main()函数返回时，程序退出，
//且程序并不等待其他goroutine（非主goroutine）结束。
func main() { //
	go Add(2, 4)
	fmt.Println(9)
	Add(2, 3)
	//time.Sleep(time.Second)
	ch := make(chan int)

	ch <- 9
	<-ch
}
func Add(x, y int) {
	z := x + y
	fmt.Println(z)
}
