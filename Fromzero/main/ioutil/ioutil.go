package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//filename := "D:\\代码\\go.cpp"
	//
	//date, err := ioutil.ReadFile(filename)
	//fmt.Println(err)
	//fmt.Println(string(date))

	//filename := "D:\\代码\\go.cpp"
	//s := "#include <stdlib.h>"
	//err := ioutil.WriteFile(filename, []byte(s), os.ModePerm)
	//
	//fmt.Println(err)

	//s := bufio.NewReader(os.Stdin)
	//dirname,_ := s.ReadString('\n')
	//fmt.Printf("%T",dirname)

	dirname := "D:\\代码"
	fileInfos,err := ioutil.ReadDir(string(dirname))
	if err != nil {
		fmt.Println(err)
		return
	}
	for i,_:=range fileInfos{
		fmt.Println(i,fileInfos[i].Name())

	}
}
