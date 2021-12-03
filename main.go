package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/nodir", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello W.")})
	http.ListenAndServe(":8082", nil)
}


