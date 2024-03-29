package main

import "fmt"

func main() { //找10000内的素数
	origin, wait := make(chan int), make(chan struct{})
	Processor(origin, wait)
	for num := 2; num < 10000; num++ {
		origin <- num
	}
	close(origin)
	<-wait
}

func Processor(seq chan int, wait chan struct{}) {
	go func() {
		prime, ok := <-seq
		if !ok {
			close(wait)
			return
		}
		fmt.Println(prime)
		out := make(chan int)
		Processor(out, wait)
		for num := range seq {
			if num%prime != 0 {
				out <- num
			}
		}
		close(out)
	}()
}
