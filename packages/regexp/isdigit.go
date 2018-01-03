package main

import "fmt"
import "regexp"

func IsDigit(user string) bool {
	pattern := "[0-9]+@xunlei\\.net"
	digit := regexp.MustCompile(pattern)
	return digit.MatchString(user)
}

func main() {
	fmt.Println("vim-go")

	if IsDigit("123456@xunlei.net") == true {
		fmt.Printf("123456@xunlei.net is digit account\n")
	}
}
