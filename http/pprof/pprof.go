package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	/*
		go func() {
			log.Println(http.ListenAndServe("localhost:8080", nil))
		}()
	*/
	log.Println(http.ListenAndServe("localhost:8080", nil))
}
