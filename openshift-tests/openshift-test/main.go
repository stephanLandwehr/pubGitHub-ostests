package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/folder", handlerFunc)
	http.ListenAndServe(":8082", nil)
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello W.")
}
