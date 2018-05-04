package main

import "fmt"
import "crypto/sha1"
import "io"

func main() {
	fmt.Println("vim-go")
	h := sha1.New()
	io.WriteString(h, "His money is twice tained:")
	io.WriteString(h, " ,taint yours and 'taint mime.")
	fmt.Printf("% X\n", h.Sum(nil))
}
