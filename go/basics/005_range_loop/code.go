package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	sum := 0
	for _, number := range numbers[1:] {
		sum += number
	}

	fmt.Println("The sum of the numbers is: %d\n", sum)
	fmt.Println(numbers)

}
