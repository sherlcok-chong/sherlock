package main

import (
	"fmt"
	"net/http"
)

func handle(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintln(w,"查询地址为",r.URL.Path)
	fmt.Fprintln(w,"查询字符串为",r.URL.RawQuery)
	fmt.Fprintln(w,"查询头中内容为",r.Header.Get("User-Agent"))
	//
	//length := r.ContentLength
	//body := make([]byte,length)
	//r.Body.Read(body)
	//fmt.Fprintln(w,"请求体为",string(body))
	r.ParseForm()
	fmt.Fprintln(w,"请求参数",r.FormValue("username"))
}

func main() {
	http.HandleFunc("/head",handle)

	http.ListenAndServe(":8080",nil)
}
