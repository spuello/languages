package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	for _, w := range strings.Fields(s) {
		count[w]++
	}
	return count
}

func main() {
	fmt.Println(WordCount("I am learning Go!"))
}
