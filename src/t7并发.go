package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Printf("%s,第%d次\n", s, i)
	}
}

func main() {
	//runtime.GOMAXPROCS(4)
	/*go say("3hello")
	go say("1world")//开一个新的Goroutines执行
	say("2hello")//当前Goroutines执行*/

	/*ci := make(chan int)
	cs := make(chan string)
	cf := make(chan interface{})*/

	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sum(a[len(a)/2:], c)
	go sum(a[:len(a)/2], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	c = make(chan int, 2) //指定channel的缓冲大小
	c <- 1

	//c <- 4
	fmt.Println(<-c)
	c <- 3
	fmt.Println(<-c)

}

//无缓冲channel
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}
