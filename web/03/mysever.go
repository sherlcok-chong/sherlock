package main

import (
	"fmt"
	"net/http"
	"time"
)

type handler struct{}

func (m *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "自我实现2.0")
}

func main() {
	hand := handler{}
	//http.Handle("/hand", &hand)
	sever := http.Server{
		Addr:        ":8080",
		Handler:     &hand,
		ReadTimeout: 20* time.Second,
	}
	sever.ListenAndServe()
}
