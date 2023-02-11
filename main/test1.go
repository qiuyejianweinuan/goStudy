package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("tom\\\\cat")
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// _表示未被使用的参数。更加规范的语言特点
// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if err != nil {
		return
	}
}
