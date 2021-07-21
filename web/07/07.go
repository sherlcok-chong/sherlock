package main

import (
	"html/template"
	"net/http"
	"web/07/contorller"
)

func IndexHandler(w http.ResponseWriter,r *http.Request)  {
	t := template.Must(template.ParseFiles("D:\\gogo\\src\\web\\07\\views\\index.html"))
	t.Execute(w,"")
}


func main() {
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/",http.StripPrefix( "/pages/",http.FileServer(http.Dir("views/pages"))))

	http.HandleFunc("/main",IndexHandler)
	http.HandleFunc("/login",contorller.Login)
	http.HandleFunc("/regist",contorller.Registered)
	http.HandleFunc("/checkUserName",contorller.CheckUsername)
	http.HandleFunc("/getBooks",contorller.GetBooks)
	http.HandleFunc("/addBook",contorller.AddBooks)
	http.HandleFunc("/deleteBook",contorller.DeleteBook)
	http.HandleFunc("/updateBook",contorller.UpdateBook)
	http.ListenAndServe(":8080",nil)

}
