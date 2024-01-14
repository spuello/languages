package main

import (
	"fmt"
)

func fibonacci() func() int {
	current := 0
	next := 1

	return func() int {
		if current == 0 {
			current, next = next, current+1
			return 0
		} else if next == 1 {
			current, next = next, current+next
			return 1
		}
		result := current + next
		current, next = next, result
		return result
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
