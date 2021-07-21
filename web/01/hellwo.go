package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"hello world",r.URL.Path)
}

func main() {
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8080",nil)
}