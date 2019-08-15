package main

import (
	"fmt"
	"time"
)

func main() {
	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	fmt.Println(arr)

	//数组可以使用另一种:=来声明
	a := [3]int{1, 2, 3}   // 声明了一个长度为3的int数组
	b := [10]int{1, 2, 3}  // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度
	fmt.Println(a, b, c)

	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	fmt.Println(doubleArray)

	slice := []int{1, 2, 3}
	fmt.Println(slice)

	//map
	//var numbers map[string] int
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	fmt.Println(numbers)
	for k, v := range numbers {
		//由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错, 在这种情况下, 可以使用_来丢弃
		//不需要的返回值
		fmt.Println(k, v)
	}
	for _, v := range numbers {
		fmt.Println(v)
	}

	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5}
	cs, ok := rating["C"]
	if ok {
		fmt.Println("csharp in map", cs)
	} else {
		fmt.Println("no csharp")
	}
	delete(rating, "C")

	//myFunc()

	/**
	Go里面switch默认相当于每个case最后带有break，匹
	配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代
	码。
	*/
	integer := 6
	switch integer {
	case 4:
		fmt.Println("integer <= 4")
	case 5:
		fmt.Println("integer <= 5")
	case 6:
		fmt.Println("integer <= 6")
		fallthrough
	case 7:
		fmt.Println("integer <= 7")
		fallthrough
	default:
		fmt.Println("default case")
	}
	fmt.Println(max(5, 8))
	fmt.Println(max(8, 5))

	args(2, 4, 6)

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//goto语句，用goto跳转到必须在当前函数内定义的标签
func myFunc() {
	i := 0
Here:
	fmt.Println(i)
	time.Sleep(1000000000)
	i++
	if i < 5 {
		goto Here
	}

}

func max(a, b int) (max int, min int) {
	if a > b {
		max, min = a, b
	} else {
		max, min = b, a
	}
	return
}

func args(arg ...int) {
	for i, j := range arg {
		fmt.Println("args", i, j)
	}
}
