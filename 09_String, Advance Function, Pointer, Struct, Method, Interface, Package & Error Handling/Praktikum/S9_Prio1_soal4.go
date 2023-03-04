package main

import (
	"fmt"
)

func getMinMax(numbers ...*int) (min int, max int) {
	if len(numbers) == 0 {
		return
	}
	min = *numbers[0]
	max = *numbers[0]
	for _, n := range numbers {
		if *n < min {
			min = *n
		} else if *n > max {
			max = *n
		}
	}
	return
}

func main() {
	var a1, a2, a3, a4, a5, a6 int
	fmt.Println("Masukkan 6 angka:")
	fmt.Scan(&a1, &a2, &a3, &a4, &a5, &a6)

	min, max := getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)

	fmt.Printf("Nilai minimum: %d\n", min)
	fmt.Printf("Nilai maksimum: %d\n", max)
}
