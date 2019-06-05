package main

import (
	"fmt"
	"runtime"
)

func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 5)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
}
