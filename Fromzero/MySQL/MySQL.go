package main

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err:= sql.Open("mysql", "root:lll2002.11.22@tcp(127.0.0.1:3306)/happy")
	if err!=nil {
		fmt.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("link fail",err)
	}
	fmt.Println("link ok")
	var uname string
	//s := "sherlock"
	s2:="lll"
	//查
	rows:= db.QueryRow("SELECT name FROM inform where ukey= ? ",s2)
	rows.Scan(&uname)
	_, err = db.Exec("insert into inform (name,ukey) value (?,?)", "happy", "l3")
	if err != nil {
		fmt.Printf("增加错误")
	}
	fmt.Printf("%v",uname)

}

func checkErr(err error){
	if err != nil {
		panic(err)
	}
}