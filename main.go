package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	randInt := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello %d", randInt)
	})
	http.ListenAndServe(":8082", nil)
	s1 := rand.NewSource(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		randInt = rand.New(s1).Intn(100)
		time.Sleep(5 * time.Second)
	}
}
