package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type User struct {
	C    chan string //
	Name string      //
	Addr string      //
}
var OnlineMap = make(map[string]User)
var Msg = make(chan string)
var db *sql.DB
var err error
var Mutex sync.Mutex
func init()  {
	db, err = sql.Open("mysql", "root:lll2002.11.22@tcp(127.0.0.1:3306)/happy")
	if err != nil {
		fmt.Println("服务器打开失败",err)
	}
}
func Find(name, key string) (uname string, now int) {
	//db,err = sql.Open("mysql", "root:lll2002.11.22@tcp(127.0.0.1:3306)/happy")
	//if err != nil {
	//	fmt.Println("连接错误", err)
	//	panic(err)
	//}
	//defer db.Close()
	//err = db.Ping()
	//if err != nil {
	//	fmt.Println("link fail",err)
	//}
	//fmt.Println("link ok")
	rows := db.QueryRow("SELECT uname,now FROM user where ukey= ? and uname = ?", key, name)
	rows.Scan(&uname, &now)
	return uname, now
}

func Insert(name, key string) bool {
	//db, err := sql.Open("mysql", "root:lll2002.11.22@tcp(127.0.0.1:3306)/happy")
	//if err != nil {
	//	fmt.Println("连接错误", err)
	//	panic(err)
	//}
	//defer db.Close()
	//err = db.Ping()
	//if err != nil {
	//	fmt.Println("link fail",err)
	//}
	var uname string
	//fmt.Println("link ok")
	rows := db.QueryRow("SELECT uname FROM user where uname = ?", name)
	rows.Scan(&uname)
	if uname == "" {
		_, err := db.Exec("insert into user (uname,ukey) value (?,?)", name, key)
		UpData(name,name,1)
		if err != nil {
			return false
		}
		return true
	}
	return false
}

func UpData(cname, name string, now int) bool{
	//db, err := sql.Open("mysql", "root:lll2002.11.22@tcp(127.0.0.1:3306)/happy")
	//if err != nil {
	//	return false
	//}
	//defer db.Close()
	_, err = db.Exec("update user set uname=?,now = ? where uname =?", cname,now, name)
	if err != nil {
		return false
	}
	return true
}
func ChangeName(uname,cname string) bool {
	//fmt.Println([]byte(cname))
	rows := db.QueryRow("SELECT uname FROM user where  uname = ?",  cname)
	var name string
	rows.Scan(&name)
	//fmt.Println(name)
	if name !=""  {
		return false
	}
	_, err = db.Exec("update user set uname=? where uname =?", cname, uname)

	if err != nil {
		return false
	}
	return true
}



