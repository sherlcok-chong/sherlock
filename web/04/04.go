package main

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"查询地址为",r.URL.Path)
	fmt.Fprintln(w,"查询字符串为",r.URL.RawQuery)
	fmt.Fprintln(w,"查询头中内容为",r.Header.Get("User-Agent"))
}

func main() {
	http.HandleFunc("/head",handle)

	http.ListenAndServe(":8080",nil)
}