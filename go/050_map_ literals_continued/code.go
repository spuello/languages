package main

import "fmt"

//If the top-level type is just a type name, you can omit it from the elements of the literal.

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}

// In Go programming language, a "type name" refers to the name given to a user-defined type.
// When you declare a custom type using the type keyword, you are creating a type name.
