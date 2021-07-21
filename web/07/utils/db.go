package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db *sql.DB
	err error
)

func init()  {
	Db,err = sql.Open("mysql","root:lll2002.11.22@tcp(127.0.0.1:3306)/happy")
	if err != nil {
		panic(err)
	}
}