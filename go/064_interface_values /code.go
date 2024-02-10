package main

import (
	"fmt"
	"math"
)

//If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

// %v: Default format for printing the value (used in debugging, general-purpose printing).

// %T: Type format for printing the type of the variable (useful in debugging and logging).

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {

}

type F float64

func (f F) M() {}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)

	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// & : Unary operator used to obtain the address of a variable.

// Example:
// i := &T{"Hello"}
// In this case, & is used to get the memory address of the variable holding a T struct.
