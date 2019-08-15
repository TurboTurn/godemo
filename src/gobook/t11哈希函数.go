package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	test := "Hi,pandaman!"
	Md5Inst := md5.New()
	Md5Inst.Write([]byte(test))
	Result := Md5Inst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	ShaInst := sha1.New()
	ShaInst.Write([]byte(test))
	Result = ShaInst.Sum([]byte(""))
	fmt.Printf("%x\n\n", Result)

	s := fmt.Sprintf("%d", 4)
	fmt.Printf(s)
	s = fmt.Sprint("hello", "yi", "shu")
	fmt.Printf(s)
}
