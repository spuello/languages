package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday ?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today:
		fmt.Println("Today")

	case today + 1:
		fmt.Println("Tomorrow")
	}
}
