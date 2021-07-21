package main

import (
	"fmt"
)


func main() {
	map1 := make(map[string]string)
	map1["姓名"] = "张三"
	map1["年龄"] = "22"
	map1["性别"] = "男"
	map1["居住地"] = "陕西"
	map2 := make(map[string]string)
	map2["姓名"] = "李四"
	map2["年龄"] = "23"
	map2["性别"] = "男"
	map2["居住地"] = "山西"
	map3 := map[string]string{"姓名": "赵六", "年龄": "20", "性别": "女", "居住地": "四川"}
	s := make([]map[string]string, 0, 3)
	s = append(s, map1)
	s = append(s, map2)
	s = append(s, map3)

	for i, j := range s {
		fmt.Println(i)
		fmt.Print(j["姓名"])
		fmt.Print(j["年龄"])
		fmt.Print(j["性别"])
		fmt.Println(j["居住地"])
	}
}
