package main

import (
	"fmt"
	"net/http"
)

func helloGo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Go")
}

func main() {
	http.HandleFunc("/api/hello", helloGo)
	http.ListenAndServe(":3000", nil)
}