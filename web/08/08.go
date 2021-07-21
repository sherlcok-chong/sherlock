package main

import (
	"html/template"
	"net/http"
)

func testIf(w http.ResponseWriter,r *http.Request)  {
	t := template.Must(template.ParseFiles("src/web/08/If.html"))
	age := 17
	t.Execute(w,age>18)
}

func main() {
	http.HandleFunc("/testIf",testIf)

	http.ListenAndServe(":8080",nil)
}
