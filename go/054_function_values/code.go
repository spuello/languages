package main

import (
	"fmt"
	"math"
)

// Functions are values too. They can be passed around just like other values.
//
// Function values may be used as function arguments and return values.

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	sum := func(x, y float64) float64 {
		return x + y
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
	fmt.Println(compute(sum))
}
