package main

import "fmt"

//The interface type that specifies zero methods is known as the empty interface

//An empty interface may hold values of any type. (Every type implements at least zero methods.)

//Empty interfaces are used by code that handles values of unknown type.
//For example, fmt.Print takes any number of arguments of type interface{}.

func main() {
	var i interface{}
	i = 42
	describe(i)

	i = "Hello"
	describe(i)

}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
