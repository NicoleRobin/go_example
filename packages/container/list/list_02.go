package main

import (
	"container/list"
	"fmt"
	"strconv"
)

type People struct {
	Name string
	Age  int
}

func main() {
	// Create a new list and put some numbers in it.
	l := list.New()
	l.PushBack(People{"zjw", 1})
	
	// Iterate through list and print its contents.
	e := l.Front()
	p, ok := (e.Value).(People)
	if ok {
		fmt.Println("Name:" + p.Name)
		fmt.Println("Age:" + strconv.Itoa(p.Age))
	} else {
		fmt.Println("e is not an People")
	}
}
