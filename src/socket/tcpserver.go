package main

import (
	"fmt"
	//"io/ioutil"
	"net"
	"os"
	"time"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":8181")
	checkErr(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkErr(err)
	for {
		fmt.Println("服务器监听：")
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		var b [512]byte
		n, err := conn.Read(b[:])
		//result, err := ioutil.ReadAll(conn)
		fmt.Println("收到消息：", string(b[0:n]))
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		fmt.Println("发送消息：", daytime)
		conn.Close()
	}

}

func checkErr(e error) {
	if e != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: s%", e.Error())
		os.Exit(1)
	}
}
