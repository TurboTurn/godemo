package main

import "os"

func main() {
	file, _ := os.Create("a.txt")
	defer file.Close()
	//file, _ = os.Open("a.txt")
	//arr :=make([]byte,1)
	file.WriteString("yishuang\n")
	file.Write([]byte("just a test1!"))

	os.Remove("a.txt")
}
