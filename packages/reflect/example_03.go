package main

import "fmt"
import "reflect"

/*
func TypeOf(i interface{}) Type
*/

func main() {
	fmt.Println("vim-go")
	var x uint8 = 'x'
	v := reflect.ValueOf(x)

	fmt.Println("type:", v.Type())
	fmt.Println("kind is uint8:", v.Kind() == reflect.Uint8)
	// x = v.Uint() // error, v.Uint() return a uint64
	x = uint8(v.Uint())
}
