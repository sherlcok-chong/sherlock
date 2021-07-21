package modle

import (
	"fmt"
	"web/07/utils"
)

type Book struct {
	ID int
	Title string
	Author string
	Price float64
	Sales int
	Stock int
	ImgPath string
}
//GetBooks 获取数据库中所有图书
func GetBooks() ([]Book,error) {
	rows,err := utils.Db.Query("select id,title,author,price,sales,stock,img_path from books")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	books := make([]Book,10)

	for rows.Next() {
		var temp Book
		err = rows.Scan(&temp.ID,&temp.Title,&temp.Author,&temp.Price,&temp.Sales,&temp.Stock,&temp.ImgPath)
		books = append(books,temp)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	return books,nil
}
func GetBook(Id int) Book {
	row := utils.Db.QueryRow("select id,title,author,price,sales,stock,img_path from books where id=?",Id)
	var temp Book
	row.Scan(&temp.ID,&temp.Title,&temp.Author,&temp.Price,&temp.Sales,&temp.Stock,&temp.ImgPath)
	return temp
}
// AddBooks 增加图书
func AddBooks(b Book) bool {
	fmt.Println(b)
	_,err := utils.Db.Exec("insert into books (title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)",b.Title,b.Author,b.Price,b.Sales,b.Stock,b.ImgPath)
	if err != nil {
		fmt.Println(err )
		return false
	}
	return true
}

func DeleteBook(ID int) bool {
	_,err:=utils.Db.Exec("delete from books where id = ?" ,ID)

	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func UpdateBook(b Book) bool {
	fmt.Println(b)
	if b.Title==""  {
		return false
	}
	_,err := utils.Db.Exec("update books set title=?,author=?,price=?,sales=?, stock=? , img_path = ? where id = ?",b.Title,b.Author,b.Price,b.Sales,b.Stock,b.ImgPath,b.ID)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
