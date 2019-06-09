package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("ys", 0777)
	os.MkdirAll("ys/y1/y2", 0777)
	err := os.Remove("ys")
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
	os.RemoveAll("ys")
}
