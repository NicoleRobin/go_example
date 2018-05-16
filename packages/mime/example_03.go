package main

import (
	"fmt"
	"mime"
)

func main() {
	fmt.Println(mime.QEncoding.Encode("utf-8", "深圳!"))
	fmt.Println(mime.QEncoding.Encode("utf-8", "Hello!"))
	fmt.Println(mime.QEncoding.Encode("UTF-8", "深圳!"))
	fmt.Println(mime.QEncoding.Encode("IOS-8859-1", "Caf\xE9"))
}
