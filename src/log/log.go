package main

import (
	"fmt"
	log1 "github.com/cihub/seelog"
	"time"
)
import "log"

func main() {
	defer log1.Flush()
	log1.Info("Hello from Seelog!")
	log.Println("geger")

	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	t := time.Now()
	fmt.Printf("当前的时间是: %d-%d-%d %d:%d:%d\n", t.Year(),
		t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	//str := t.Format("2006-01-02 15:04:05")
	//str := t.Format("2006-01-02")
	//str := t.Format("15:04:05")
	str := t.Format("2006")
	fmt.Println(str)
	fmt.Println("over")
	time.Sleep(time.Second * 2)
	go func() {
		fmt.Println("sleep")
		time.Sleep(time.Second)
		fmt.Println("sleep2")
	}()

}
