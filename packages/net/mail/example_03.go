package main

import (
	"fmt"
	"net/mail"
)

func main() {
	addr := mail.Address{"zjw", "lit050528@gmail.com"}
	fmt.Println(addr.String())
}
