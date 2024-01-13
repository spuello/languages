package main

import "fmt"

//import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var s [][]uint8

	for i := 0; i < dy; i++ {
		var a []uint8
		for j := 0; j < dx; j++ {
			a = append(a, uint8((i+j)/2))

		}
		s = append(s, a)
	}
	return s
}

func main() {
	fmt.Println(Pic(3, 6))
	//pic.Show(Pic)
}
