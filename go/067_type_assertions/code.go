package main

import "fmt"

// A type assertion provides access to an interface value's underlying concrete value.

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)
	s, ok := i.(string)
	fmt.Println(s, ok)

	//Zero-value and false
	f, ok := i.(float64)
	fmt.Println(f, ok)
	
	//Panic
	f = i.(float64)
	fmt.Println(f)
}
