package main

import (
	"net/http"
	"text/template"
)

func testTemplate(w http.ResponseWriter,r *http.Request)  {
	t,_:=template.ParseFiles("05.html")

	t.Execute(w,"????")
}

func main() {
	http.HandleFunc("/testTemplate",testTemplate)


	http.ListenAndServe(":8088",nil)
}
