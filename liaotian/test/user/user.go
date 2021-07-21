//client
package main

import (
	"fmt"
	"net"
	"os"
)

var ch chan int = make(chan int)
var nickname string

func reader(conn net.Conn) {
	buff := make([]byte, 256)
	for {
		j, err := conn.Read(buff)
		if err != nil {
			ch <- 1
			break
		}
		fmt.Printf("%s\n", buff[0:j])
	}
}
func main() {
	//if len(os.Args) != 2 {
	//	fmt.Fprintf(os.Stderr, "Usage:%s  host:port", os.Args[0])
	//	os.Exit(1)
	//}
	//service := os.Args[1]
	//TcpAdd, _ := net.ResolveTCPAddr("tcp", service)
	//TcpAdd, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	conn, err := net.Dial("tcp", "127.0.0.1:1122")
	if err != nil {
		fmt.Println("服务没打开")
		os.Exit(1)
	}
	defer conn.Close()
	go reader(conn)
	fmt.Println("请输入昵称")
	fmt.Scanln(&nickname)
	fmt.Println("你的昵称是：", nickname)
	//var str string
	for {
		var msg string
		fmt.Scan(&msg)
		fmt.Print("<" + nickname + ">" + "说:")
		//for i, _ := range msg {
		//	fmt.Printf("%c", msg[i])
		//}
		fmt.Println(msg)
		b := []byte("<" + nickname + ">" + "说：" + msg)
		conn.Write(b)
		select {
		case <-ch:
			fmt.Println("server发生错误，请重新连接")
			os.Exit(2)
		default:
		}
	}
}
