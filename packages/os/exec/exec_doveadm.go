package main

import "fmt"

// import "os"
import "os/exec"
import "log"
import "bytes"
import "strings"

func doveadm(user string) {
	cmd := exec.Command("/usr/local/bin/doveadm", "mailbox", "status", "-u", user, "-t", "unseen", "INBOX")
	// fmt.Printf("%s\n", cmd.Path)
	// fmt.Printf("%v\n", cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}

	result := out.String()
	// fmt.Printf("User:%s Result:%s\n", user, result)

	values := strings.Split(result, "=")
	if len(values) >= 2 {
		fmt.Printf("User:%s unseen:%s\n", user, values[1])
	}
}

func main() {
	doveadm("1517825819@xunlei.net")
	doveadm("151782581@xunlei.net")
}
