package view

import (
	"fmt"
	"liaotian/liaotian/model"
	"os"
)

func OpenView() (choice int) {
	fmt.Println("1.登录")
	fmt.Println("2.注册")
	fmt.Scanf("%d\n", &choice)
	if choice != 1 && choice != 2 {
		fmt.Println("选择错误")
		return 0
	}
	return choice
}
func Select(choice int) (bool, string) {

	var name, key string
	fmt.Printf("用户名：")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("密码：")
	fmt.Scanf("%s\n", &key)
	if choice == 1 {
		return LogIn(name, key), name
	}
	return register(name, key), name
}

func LogIn(name, key string) bool {

	//fmt.Println(name, key)
	//fmt.Println("输入查询成功")
	key=Hash([]byte(key),len(key))
	uname, now := model.Find(name, key)
	if uname == "" {
		fmt.Println("用户名或密码错误")
		return false
	} else if now == 1 {
		fmt.Println("该用户已登录")
		return false
	}
	return true
}
func register(name, key string) bool {
	key=Hash([]byte(key),len(key))
	ok := model.Insert(name, key)
	if ok == true {
		fmt.Println("注册成功")
		return true
	}
	fmt.Println("注册失败")
	return false
}

func Hash(name []byte,n int) string {
	for i, val := range name {
		name[i] = val % 10+111
	}
	return string(name[:n])
}
func ChangeMyName(name string) (string,  bool) {
	fmt.Println("是否修改姓名(y)")
	s := make([]byte, 20)
	n, err := os.Stdin.Read(s)
	if err != nil {
		fmt.Println("输入错误")
		return "",false
	}
	if string(s[:n-2]) != "y" {
		return "",false
	}
	fmt.Println("输入新昵称")
	n, err = os.Stdin.Read(s)
	if err != nil {
		fmt.Println("输入错误")
		return "",false
	}
	re:=model.ChangeName(name,string(s[:n-2]))
	if !re {
		fmt.Println("修改失败")
	}
	fmt.Println("修改成功")
	return string(s[:n-1]),true
}