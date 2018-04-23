package main

import "fmt"
import "reflect"

/*
func TypeOf(i interface{}) Type
*/

func main() {
	fmt.Println("vim-go")
	var x float64 = 3.4
	v := reflect.ValueOf(x)

	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}