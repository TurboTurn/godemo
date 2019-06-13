package main

import (
	"fmt"
	"log"
)

func main() {
	a := []int{1, 2, 3} //切片或数组 [3]int{1,2,3}
	b := a
	b[1]++ //如果是切片，会改变a；如果是数组不会改变a
	fmt.Println(a, b)

	r := Rect{1, 2, 3, 4}
	r1 := new(Rect) //new 返回的指针
	r2 := &Rect{1, 2, 3, 4}
	r3 := NewRect(1, 2, 3, 4)
	fmt.Println(r.Area())
	fmt.Println(r2.Area())
	fmt.Println(r, r1, r2, r3)

	/*job := new(Job)
	job.Start()*/
	defer func() {
		if r := recover(); r != nil {

		}
	}()

}

type Rect struct {
	x, y          float64
	width, height float64
}

func NewRect(x float64, y float64, width float64, height float64) *Rect {
	return &Rect{x: x, y: y, width: width, height: height}
}
func (r *Rect) Area() float64 {
	return r.width * r.height
}

type Job struct {
	Command string
	*log.Logger
}

func (job *Job) Start() {
	job.Println("starting now...")
}
