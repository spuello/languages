package main

import "fmt"

//Init statement; condition expression; post statement

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("The sum of the numbers is: %d\n\n", sum)

}
