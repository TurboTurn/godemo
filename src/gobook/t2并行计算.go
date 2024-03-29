package main

import "fmt"

func sum(values []int, result chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	result <- sum
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := make(chan int, 2)
	go sum(values[:len(values)/2], result)
	go sum(values[len(values)/2:], result)
	sum1, sum2 := <-result, <-result
	fmt.Println(sum1 + sum2)
}
