package main

import "fmt"

func main() {
	slice := make([]int, 5) //创建一个初始元素个数为5的数组切片，元素初始值为0
	fmt.Println(slice)

	//创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间len=5,cap=10
	slice1 := make([]int, 5, 10)
	fmt.Println(slice1)

	//直接创建并初始化包含5个元素的数组切片,事实上还会有一个匿名数组被创建出来，只是不需要我们来操心而已
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice3)

	m := make(map[string]string, 10)
	m["s"] = "s"
	a, ok := m["s"]
	if ok {
		fmt.Println(a)
	}
	delete(m, "s")
}
