package main

import (
	"fmt"
)

func main() {
	var v5 struct {
		f int
	}
	//v5.f = 5

	v6 := "fdf"
	v7 := &v6
	fmt.Println(v5)
	fmt.Println(v7)
	fmt.Println(*v7)
	fmt.Println(&v7)
	const home = 56
	var f1 float32
	var f2 float32 = 6
	if f1 == f2 {
		fmt.Println(f1)
	}
	fmt.Println(^-3)
	value1 := 3.4 + 5i
	fmt.Println(value1)

	str := "Hello,世界"
	n := len(str)
	for i := 0; i < n; i++ {
		ch := str[i] // 依据下标取字符串中的字符，类型为byte
		fmt.Println(i, ch)
	}
	for i, ch := range str {
		fmt.Println(i, ch) //ch的类型为rune
	}
}
