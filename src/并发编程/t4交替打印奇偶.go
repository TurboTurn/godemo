package main

import "fmt"

func main() {
	c := make(chan int, 100)
	i := 1
	b := &i
	go t1(*b, c)
	t2(*b, c)
}

func t1(i int, c chan int) {
	for ; i <= 103; i++ {
		if i&1 == 1 {
			fmt.Println("t1生产", i)
			c <- i
		} else {
			fmt.Println("t1消费", <-c)
		}
	}
}
func t2(i int, c chan int) {
	for ; i <= 100; i++ {
		if i&1 == 0 {
			fmt.Println("t2生产", i)
			c <- i
			if i == 100 {
				i--
				break
			}
		} else {
			fmt.Println("t2消费", <-c)
		}
	}
}
