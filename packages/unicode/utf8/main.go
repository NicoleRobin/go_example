package main

import "fmt"
import "unicode/utf8"

func main() {
	fmt.Println("vim-go")
	b := []byte("hello, 世界")
	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b)
		fmt.Printf("%c %v\n", r, size)
		b = b[:len(b)-size]
	}
}
