package main

import "fmt"

func main() {
	fmt.Println("good")
	r1 := Rectangle{3, 5}
	fmt.Println(r1.area())

	var a interface{} = "hello"
	fmt.Println(a)
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}
