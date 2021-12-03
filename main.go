package main

import (
	"fmt"
	"net/http"
	main2 "https://github.com/stephanLandwehr/os-test/openshift-tests/openshift-test/main2.go"
)

func main() {
	http.HandleFunc("/nodir", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello W.")})
	main2.main()
}


