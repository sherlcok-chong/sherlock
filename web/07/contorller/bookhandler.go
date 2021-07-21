package contorller

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"web/07/modle"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := modle.GetBooks()
	//fmt.Println(books)
	if err != nil {
		fmt.Println(err)
	}
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	t.Execute(w, books)
}
func AddBooks(w http.ResponseWriter, r *http.Request) {
	var b modle.Book
	b.Title = r.PostFormValue("title")
	b.Price, _ = strconv.ParseFloat(r.PostFormValue("price"), 32)

	b.Author = r.PostFormValue("author")

	temp, _ := strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
	b.Sales = int(temp)
	temp, _ = strconv.ParseInt(r.PostFormValue("stock"), 10, 0)
	b.Stock = int(temp)
	b.ImgPath = "/static/img/default.jpg"
	fmt.Printf("%T%T%T%T%T", b.Price, b.Title, b.Sales, b.Author, b.Stock)
	re := modle.AddBooks(b)
	if re != true {
		fmt.Println("添加失败")
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	temp, _ := strconv.ParseInt(bookId, 10, 0)
	re := modle.DeleteBook(int(temp))
	if re != true {
		fmt.Println("删除失败")
	} else {
		fmt.Println("删除成功")
	}

}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/pages/manager/book_modify.html"))
	bookId := r.FormValue("bookId")
	temp, _ := strconv.ParseInt(bookId, 10, 0)
	book := modle.GetBook(int(temp))
	fmt.Println(book)
	t.Execute(w,book)

	book.Title = r.PostFormValue("title")
	book.Price, _ = strconv.ParseFloat(r.PostFormValue("price"), 32)

	book.Author = r.PostFormValue("author")

	temp, _ = strconv.ParseInt(r.PostFormValue("sales"), 10, 0)
	book.Sales = int(temp)
	temp, _ = strconv.ParseInt(r.PostFormValue("stock"), 10, 0)
	book.Stock = int(temp)
	book.ImgPath = "/static/img/default.jpg"
	re := modle.UpdateBook(book)
	if re != true {
		fmt.Println("修改失败")
	}
}
