package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`[
	{"Name":"Platpus", "Order":"Monotremata"},
	{"Name":"Quoll", "Order":"Dasyuromorphid"}
	]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals)
}
