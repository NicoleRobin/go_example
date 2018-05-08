package main

import "fmt"
import "crypto/sha1"
import "io"

func main() {
	fmt.Println("vim-go")
	h := sha1.New()
	io.WriteString(h, "hello")
	fmt.Printf("% X\n", h.Sum(nil))
}
