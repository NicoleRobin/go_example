package main

import "fmt"
import "crypto/sha1"
import "os"
import "log"
import "strconv"

func main() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Printf("Usage:%s file length\n", os.Args[0])
		return
	}

	var err error
	var length int = 102400000
	if len(os.Args) == 3 {
		length, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	buffer := make([]byte, length)
	ret, err := file.Read(buffer)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("read bytes:%d\n", ret)

	fmt.Printf("%X\n", sha1.Sum(buffer))

	file.Close()
}
