package main

import (
	"fmt"
	"runtime"
)

func main() {
	//获取go安装目录
	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.NumCPU())
}
