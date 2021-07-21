package main

import (
	"fmt"
	"liaotian/liaotian/model"
	"net"
)

func main() {
	//choice := view.OpenView()
	listener, err := net.Listen("tcp", "127.0.0.1:1122")
	if err != nil {
		fmt.Printf("listen err -> %v", err)
		return
	}
	defer listener.Close()
	//给每一个用户发送上线消息

	//阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept err -> %v", err)
			continue
		}
		//name :=
		fmt.Println(conn.RemoteAddr().String(), "link ok")
		go HandleConn(conn)
	}
}

// HandleConn 处理用户连接
func HandleConn(conn net.Conn) {
	defer conn.Close()
	CliAddr := conn.RemoteAddr().String()
	buf := make([]byte, 2048)
	n, _ := conn.Read(buf)
	cli := model.User{C: make(chan string), Name: string(buf[:n]), Addr: CliAddr}
	//把结构体加到map
	model.OnlineMap[CliAddr] = cli
	//广播在线转发消息
	go Manager()
	for _, val := range model.OnlineMap {
		fmt.Println(val)
	}
	model.Msg <- cli.Name + "上线了"
	//新开一个协程，专门给客户端发消息
	go WriteMsgToClient(cli, conn)

	func() {

		for {

			n, _ := conn.Read(buf)
			if n == 0 {
				go Manager()
				model.Msg <- cli.Name + "下线了"
				//fmt.Printf("%s下线了",cli.Name)
				delete(model.OnlineMap, CliAddr)
				model.UpData(cli.Name, cli.Name, 0)
				return
			} else {
				fmt.Println("收到消息")
			}

			//转发此内容
			msg := string(buf[:n-1])
			if len(msg) == 3 && msg == "who" {
				conn.Write([]byte("用户列表：\n"))
				for _, val := range model.OnlineMap {
					conn.Write([]byte(val.Addr + ":" + val.Name + "\n"))
				}
			}
			if  msg == "change name" {
				//fmt.Println(len(model.OnlineMap))
				n , _ := conn.Read(buf)
				cname := string(buf[:n-1])
				//fmt.Println(cname)
				temp := model.OnlineMap[CliAddr]
				temp.Name=cname
				model.OnlineMap[CliAddr]=temp
				//fmt.Println(model.OnlineMap)
				cli.Name=cname
				continue
			}
			go Manager()
			model.Msg <- MakeMassage(cli, msg)
		}
	}()
}
func MakeMassage(cli model.User, msg string) (buf string) {
	buf = "[" + cli.Name + "]" + ":" + msg
	return
}
func Manager() {
	msg := <-model.Msg
	fmt.Println("接收到了消息")
	for _, cli := range model.OnlineMap {
		cli.C <- msg
	}
}
func WriteMsgToClient(cli model.User, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

