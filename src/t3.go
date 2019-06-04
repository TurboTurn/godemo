package main

import "fmt"

func main()  {
	/*var P person
	P.name = "Astaxie"
	P.age = 25
	fmt.Println(P)*/

	//按照顺序提供初始化值
	/*P := person{"tom",30}
	fmt.Println(P)*/

	//通过field:value的方式初始化，这样可以任意顺序
	/*P := person{age:28, name:"yishuang"}
	fmt.Println(P)*/

	stu := student{human{"ys",23,63},human2{4},""}
	stu.human.age = 22
	fmt.Println(stu)
}

type person struct {
	name string
	age int
}

type human struct {
	name string
	age int
	weight int
}
type human2 struct {
	age int
}

type student struct {
	human
	human2
	speciality string
}

func (s student)String() string  {//调用fmt.print会输出此处
	return "rr✆"
}



