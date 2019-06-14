package main

import (
	"fmt"
	"time"
)

func main() {
	prod := make(chan int)
	go producer(prod)
	consumer(prod)
	fmt.Println("comsumer")
}

func producer(prod chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Nanosecond)
		fmt.Println("生产：", i)
		prod <- i
	}
}
func consumer(cons chan int) {
	for i := 0; i < 10; i++ {
		value := <-cons
		fmt.Println("消费：", value)
	}

}
