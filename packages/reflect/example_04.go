package main

import "fmt"
import "reflect"

/*
func TypeOf(i interface{}) Type
*/

func main() {
	fmt.Println("vim-go")
	type MyInt int
	var x MyInt = 7
	v := reflect.ValueOf(x)

	fmt.Println("type:", v.Type())
	// Kind() return its underlying type, not its static type
	// MyInt is its static type
	fmt.Println("kind is int:", v.Kind() == reflect.Int)
}
