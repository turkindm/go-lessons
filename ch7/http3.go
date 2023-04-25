package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Index")
}

func main() {
	http.HandleFunc("/index", index)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
