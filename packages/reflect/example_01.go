package main

import "fmt"
import "reflect"

/*
func TypeOf(i interface{}) Type
*/

func main() {
	fmt.Println("vim-go")
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	fmt.Println("value:", reflect.ValueOf(x))
	fmt.Println("value:", reflect.ValueOf(x).String())
}
