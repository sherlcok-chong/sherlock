package modle

import (
	"fmt"
	"web/07/utils"
)

type Users struct {
	ID int
	Username string
	Password string
	Email string
}

// CheckUserName 查询用户名和密码
func CheckUserName(username string) (Users ,error){
	fmt.Println("数据库了",username)
	row := utils.Db.QueryRow("select id,username,password,email from users where username=? ",username)

	user := Users{}
	err:=row.Scan(&user.ID,&user.Username,&user.Password,&user.Email)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("数据库了",user)
	return user ,nil
}

func CheckPassword(username string, password string)(Users, bool ){
	user,err:= CheckUserName(username)
	if err != nil {
		panic(err)
	}
	if user.Password!=password||password=="" {
		return user,false
	}
	return user,true
}

func InsertUser(user Users) bool {
	_,err := utils.Db.Exec("insert into users (id,username,password,email) value (?,?,?,?) ",user.ID,user.Username,user.Password,user.Email)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}