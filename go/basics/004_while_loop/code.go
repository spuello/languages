package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	i := 0

	for i < len(numbers) {
		sum += numbers[i]
		i++
	}
	fmt.Println("The sum is : %d\n", sum)
}
