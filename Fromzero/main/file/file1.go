package main

import (
	"fmt"
	"os"
)

func main() {
	//fileinfo, err := os.Stat("D:\\代码\\0000.cpp")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("%T\n", fileinfo)
	//
	//fmt.Println(fileinfo.Name())
	//
	//fmt.Println(fileinfo.ModTime())
	//
	//fmt.Println(fileinfo.Size())
	//fmt.Println(fileinfo.Mode())

	//err2 := os.Mkdir("D:\\代码\\创建", os.ModePerm)
	//if err != nil {
	//	fmt.Println(err2)
	//	return
	//}
	//fmt.Println("成功")

	//err2 := os.MkdirAll("D:\\代码\\创建\\345\\3453\\345", os.ModePerm)
	//if err2 != nil {
	//	fmt.Println(err2)
	//	return
	//}
	//file1,err:=os.Create("D:\\代码\\go.cpp")
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//fmt.Println(file1)
	//删目录
	err := os.Remove("D:\\代码\\0000.exe")
	if err!=nil {
		fmt.Println(err)
		return
	}
	//删除目录只能删除空目录

}
