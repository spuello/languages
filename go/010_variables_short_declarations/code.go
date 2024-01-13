package main

import "fmt"

// This declaration only works inside a function
func main() {
	var i, j int = 1, 3
	c, python, java := true, false, "no"

	fmt.Println(i, j, c, python, java)
}
