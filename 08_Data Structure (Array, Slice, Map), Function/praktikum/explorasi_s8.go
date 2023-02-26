package main

import (
	"fmt"
	"math"
)

func main() {
	// inisialisasi matriks
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}

	// hitung jumlah diagonal kiri ke kanan
	sum1 := 0
	for i := 0; i < len(matrix); i++ {
		sum1 += matrix[i][i]
	}

	// hitung jumlah diagonal kanan ke kiri
	sum2 := 0
	for i := 0; i < len(matrix); i++ {
		sum2 += matrix[i][len(matrix)-1-i]
	}

	// hitung selisih absolut
	diff := int(math.Abs(float64(sum1 - sum2)))

	// tampilkan hasil
	fmt.Println("Matriks:")
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Printf("Selisih absolut diagonal: %d\n", diff)
}
