package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello inner!")
	})
	http.HandleFunc("/hello/hello2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello inner2!")
	})
	addr := "127.0.0.1:8080"
	fmt.Println("starting server at port 8080")
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
