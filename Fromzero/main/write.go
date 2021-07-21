package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "D:\\代码\\go.cpp"
	file,err := os.OpenFile(filename,os.O_WRONLY|os.O_APPEND,os.ModePerm)

	if err!= nil{
		fmt.Println(err)
		return
	}

	defer file.Close()
	n,err2:=file.WriteString("int main(){\n}")
	fmt.Println(n,err2)
}
