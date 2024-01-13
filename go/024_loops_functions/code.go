package main

import "fmt"

func main() {
	Sqrt(100)

}
func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("The square root is %g\n", z)
	}

	return z
}
