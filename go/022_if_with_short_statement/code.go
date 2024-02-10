package main

import (
	"fmt"
	"math"
)

//Like for, the if statement can start with a short statement to execute before the condition.

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 100, 20),
	)
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
