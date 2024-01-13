package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Samuel Jose", "Puello")
	fmt.Println(a, b)
}
