package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	html, _ := ioutil.ReadFile("hello.html")
	_, _ = fmt.Fprintln(w, string(html))
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server failed, err: %v \n", err)
		return
	}
}
