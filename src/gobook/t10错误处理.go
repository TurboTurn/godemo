package main

import (
	"errors"
	"fmt"
)

func main() {
	var e error
	fmt.Println(e)
	e = errors.New("errtrg")
	fmt.Println(e)
	fmt.Println(e.Error())

	e = &MyError{"myError"}
	fmt.Println(e)

	//无论foo()中是否触发了错误处理流程，该匿名defer函数都将在函数退出时得到执行
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover()执行", r)
		}
	}()
	panic(MyError{"模拟发生错误"})
}

type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return e.message
}
