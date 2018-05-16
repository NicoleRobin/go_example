package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/mail"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage:%s mail\n", os.Args[0])
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	m, err := mail.ReadMessage(file)
	if err != nil {
		log.Fatal(err)
	}

	headers := m.Header
	for k, v := range headers {
		fmt.Println(k, v)
	}

	body, err := ioutil.ReadAll(m.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("==================================================")
	fmt.Printf("%s", body)
}
