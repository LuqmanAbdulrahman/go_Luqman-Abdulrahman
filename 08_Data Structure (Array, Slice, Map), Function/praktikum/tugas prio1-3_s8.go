package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	countMap := make(map[int]int)
	var uniqueSlice []int

	// iterasi string input
	for _, char := range angka {
		num, err := strconv.Atoi(string(char))
		if err == nil {
			// tambahkan jumlah kemunculan tiap angka pada map
			countMap[num]++
		}
	}

	// ambil angka yang hanya muncul sekali pada input
	for num, count := range countMap {
		if count == 1 {
			uniqueSlice = append(uniqueSlice, num)
		}
	}

	return uniqueSlice
}

func main() {
	fmt.Println(munculSekali("1234123"))    // [4]
	fmt.Println(munculSekali("76523752"))   // [6,3]
	fmt.Println(munculSekali("12345"))      // [1 2 3 4 5]
	fmt.Println(munculSekali("1122334455")) // []
	fmt.Println(munculSekali("0872504"))    // [8 7 2 5 4]
}
