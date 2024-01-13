package main

import "fmt"

//A slice does not store any data, it just describes a section of an underlying array.

func main() {

	names := [4]string{"Bach", "Chopin", "Cervantes", "García"}

	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

}
