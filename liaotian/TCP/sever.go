package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:1122")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	//接收多个请求
	for  {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go HandleConn(conn)

	}
}

// HandleConn 处理用户请求
func HandleConn(conn net.Conn)  {
	defer conn.Close()

	addr := conn.RemoteAddr().String()
	fmt.Println(addr,"link ok")

	buf := make([]byte, 2048)
	for {
		n, err1 := conn.Read(buf)
		if err1 != nil {
			fmt.Println("err1", err1)
			return
		}
		fmt.Printf("%s\n", string(buf[:n]))
		if string(buf[:n]) == "exit" {
			fmt.Println("结束了")
			return
		}
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}