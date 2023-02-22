package main

import "fmt"

func main() {
	segitiga()
}
func segitiga() {
	for x := 1; x <= 5; x++ {
		for y := 5; y >= x; y-- {
			fmt.Print(" ")
		}
		for z := 1; z <= x; z++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
