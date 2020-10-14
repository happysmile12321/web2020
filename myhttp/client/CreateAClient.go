package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	url := "http://127.0.0.1:10080"
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	//Get1
	//_, err := client.Get(url)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//Get2
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("date", `date: Sat, 03 Oct 2020 00:08:23 GMT`)
	req.Header.Add("If-none-Match", `W/wyzzy`)
	req.Header.Add("content-type", `text`)
	resp, _ := client.Do(req)
	buffer := make([]byte, 4)
	var html = []string{}
	for {
		_, err := resp.Body.Read(buffer)
		html = append(html, string(buffer))
		if err != nil {
			if err == io.EOF {
				fmt.Println("read end!!!")
				break
			}
		}
	}
	fmt.Println(html)
}

func redirectPolicyFunc(*http.Request, []*http.Request) error {
	return nil
}
