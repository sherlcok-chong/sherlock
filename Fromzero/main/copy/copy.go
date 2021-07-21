package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	srcFile := "D:\\代码\\go.cpp"
	destFile := "D:\\代码\\go2.cpp"
	total, err := CopyFile1(srcFile, destFile)
	fmt.Println(total, err)
}
func CopyFile2(srcFile, destFile string) (int64, error) {
	file1, err1 := os.Open(srcFile)
	if err1 != nil {
		return 0,err1
	}
	file2,err2 :=os.OpenFile(destFile,os.O_CREATE|os.O_WRONLY,os.ModePerm)
	if err2 !=nil{
		return 0, err2
	}
	defer file1.Close()
	defer file2.Close()

	return io.Copy(file2,file1)
}
func CopyFile1(srcFile, destFile string) (int, error) {
	file1, err1 := os.Open(srcFile)
	if err1 != nil {
		return 0, err1
	}
	file2, err2 := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err2 != nil {
		return 0, err2
	}
	defer file1.Close()
	defer file2.Close()
	bs := make([]byte, 1024, 1024)
	n := -1
	total := 0
	for {
		n, err1 = file1.Read(bs)
		if err1 == io.EOF || n == 0 {
			fmt.Println("完了")
			break
		} else if err1 != nil {
			fmt.Println("错误")
			return total, err1
		}
		total += n
		file2.Write(bs[:n])
	}

	return total, nil
}
