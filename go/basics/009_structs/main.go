package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	chopin := Person{firstName: "Frederik", lastName: "Chopin"}

	fmt.Printf("%+v", chopin)
}
