package main

import (
	"fmt"
)

type lessadd interface {
	less(b Integer) bool
	add(b Integer)
}

type Integer int

func (i Integer) less(b Integer) bool {
	return i < b
}
func (a Integer) add(b Integer) { //不会改变
	a += b
}
func (a *Integer) add1(b Integer) { //会改变a
	*a += b
}
func main() {
	var a Integer = 1
	if a.less(5) {
		fmt.Println(a)
	}
	a.add(3)
	fmt.Println(a)
	a.add1(3)
	fmt.Println(a)

	var b lessadd = &a
	b.add(4)
	f, ok := b.(Integer)
	if ok {
		fmt.Println(f)
	}
}
