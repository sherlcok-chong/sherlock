package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "D:\\代码\\go.cpp"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	bs := make([]byte, 1024, 1024)

	//n, err2 := file.Read(bs)
	//
	//fmt.Println(err2)
	//fmt.Println(n)
	//fmt.Println(string(bs))
	//
	//n, err2 = file.Read(bs)
	//
	//fmt.Println(err2)
	//fmt.Println(n)
	//fmt.Println(string(bs))
	//
	//n, err2 = file.Read(bs)
	//
	//fmt.Println(err2)
	//fmt.Println(n)
	//fmt.Println(string(bs))
	n :=-1
	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println(123)
			break
		}
		fmt.Println(string(bs[0:n]))
	}
}
