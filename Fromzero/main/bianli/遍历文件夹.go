package main

import (
	"fmt"
	"io/ioutil"
	"time"
)
var cnt2 int
func main() {
	start := time.Now()
	dirname := "C:\\"
	listFiles(dirname)
	ended := time.Now().Sub(start)
	fmt.Println(ended,cnt2)
}

func listFiles(dirname string) {
	fileInfos, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _,i :=range fileInfos{
		filename := dirname+ "/" +i.Name()
		//fmt.Println(filename)
		cnt2++
		if i.IsDir(){
			listFiles(filename)
		}
	}
	return
}
