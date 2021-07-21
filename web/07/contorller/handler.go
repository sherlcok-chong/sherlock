package contorller

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"web/07/modle"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	//调用测验函数
	fmt.Println(username, password)
	user, re := modle.CheckPassword(username, password)
	fmt.Println(user)
	if re != false {
		//正确
		fmt.Println("查询成功")
		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, "")
	} else {
		//false
		fmt.Println("查询错误")
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}
}
func Registered(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	user, re := modle.CheckPassword(username, password)
	fmt.Println(user)
	if re != false {
		//正确
		fmt.Println("增加失败")
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "")
	} else {
		//false
		fmt.Println("增加成功")
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		modle.InsertUser(modle.Users{ID: 1122, Username: username, Password: password, Email: email})
		t.Execute(w, "")
	}
	rand.Int()
}
func CheckUsername(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	fmt.Println("传来的参数是：")
	user, err := modle.CheckUserName(username)
	if err != nil {
		fmt.Println(err)
	}
	if user.Username == "" {
		w.Write([]byte("用户名可用"))
	} else {
		w.Write([]byte("用户名已存在"))
	}
}
