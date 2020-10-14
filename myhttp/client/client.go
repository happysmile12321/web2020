package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	url := "http://127.0.0.1:10080"
	json := `
"people":[
"zhangsan",
"lisi",
"wangwu"
]
	`
	resp, err := http.Post(url, "text", strings.NewReader(json))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

}
