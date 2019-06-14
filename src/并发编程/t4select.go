package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case ch <- 3:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println("Value received:", i)
	}
}
