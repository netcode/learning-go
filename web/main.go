package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello GOWeb")
}

func main() {
	fmt.Println("hello web")
	http.HandleFunc("/", index)
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong")
	})
	http.ListenAndServe(":3003", nil)
}
