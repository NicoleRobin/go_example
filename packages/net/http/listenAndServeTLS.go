package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("About to listen on 8080. Go to https://10.10.33.35:8080/")
	err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	log.Fatal(err)
}
