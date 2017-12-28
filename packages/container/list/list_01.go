package main

import (
		"container/list"
		"fmt"
	   )

func main() {
	// Create a new list and put some numbers in it
l := list.New()
	   e4 := l.PushBack(4)
	   e1 := l.PushBack(1)
	   l.InsertBefore(3, e4)
	   l.InsertAfter(2, e1)

	   // Iterate though list and print its contents
	   for e := l.Front(); e != nil; e = e.Next() {
		   fmt.Println(e.Value)
	   }
}
