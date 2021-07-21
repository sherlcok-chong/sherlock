package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type person struct {
	name string
	id int
}

func testRange(w http.ResponseWriter,r *http.Request)  {
	t := make([]person,10)
	t[0] = person{"123",456}
	t[1] = person{name: "412",id: 789}

	temp := template.Must(template.ParseFiles("src/web/09/range.html"))
	fmt.Println(t)
	temp.Execute(w,t)
}

func main() {
	http.HandleFunc("/range",testRange)

	http.ListenAndServe(":8080",nil)
}