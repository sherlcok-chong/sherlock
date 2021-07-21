package main

import (
	"fmt"
	"liaotian/liaotian/model"
	"liaotian/liaotian/view"
	"net"
	"os"
)

func main() {
	//做出选择注册或者登录

	var uname string
	for {
		choice := view.OpenView()
		for choice == 0 {
			choice = view.OpenView()
		}
		ok, name := view.Select(choice)
		if ok == true {
			uname = name
			break
		}
	}
	model.UpData(uname, uname, 1)
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:1122")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	conn.Write([]byte(uname))
	fmt.Println(conn.RemoteAddr().String())
	//接收数据
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("读取err", err)
				return
			}
			if string(buf[:n]) == "exit" {
				return
			}
			fmt.Printf("%s", string(buf[:n]))
		}
	}()

	//发送数据
	s := make([]byte, 2048)
	for {

		n, err := os.Stdin.Read(s)
		if err != nil {
			fmt.Printf("键盘%v", err)
			return
		}
		//fmt.Printf("%v,%v",string(s[:n]),len(s))
		if string(s[:n-2]) == "exit" {
			return
		}
		if string(s[:n-2]) == "change name" {
			cname, ans := view.ChangeMyName(uname)
			//fmt.Println(ans)
			if ans == true {
				conn.Write(s[:n-1])
				conn.Write([]byte(cname))
				continue
			}
		}
		conn.Write(s[:n-1])
	}
}












