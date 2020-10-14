package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello,Go1")
	buffer := make([]byte, 2)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("read done!")
				break
			} else {
				log.Fatal(err)
			}
		}
		fmt.Println("read: ", n, " bytes,and buffer is :", string(buffer))
	}

}
