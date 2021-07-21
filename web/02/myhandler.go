package main

import (
	"fmt"
	"net/http"
)

type handler struct{}

func (m *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "自我实现")
}

func main() {
	hand := handler{}
	http.Handle("/hand", &hand)
	http.ListenAndServe(":8080", nil)
}
