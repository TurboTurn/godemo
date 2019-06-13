package main

import (
	"fmt"
)

func main() {
	myFun(3, 45, 6, 345)
}
func myFun(args ...int) {
	fmt.Println(args)
	myFun1(args[1:]...)
	myFun2(1, 5, 8.3, 22)
}

func myFun1(args ...int) {
	fmt.Println(args)
}

func myFun2(args ...interface{}) {
	for i, v := range args {
		fmt.Println(i, v)
	}
	fmt.Println(args)
}
