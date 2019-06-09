package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":8181")
	checkErr(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkErr(err)
	_, err = conn.Write([]byte("hello"))
	fmt.Println("发送消息")
	checkErr(err)
	result, err := ioutil.ReadAll(conn)
	checkErr(err)
	fmt.Print("收到回复：")
	fmt.Println(string(result))
	conn.Close()
}
func checkErr(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: s%", e.Error())
		os.Exit(1)
	}
}
