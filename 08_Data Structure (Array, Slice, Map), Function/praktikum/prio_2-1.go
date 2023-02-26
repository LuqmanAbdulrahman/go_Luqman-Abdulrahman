package main

import "fmt"

func PairSum(arr []int, target int) []int {
	// inisialisasi variabel map
	numsMap := make(map[int]int)

	// iterasi array dan cari pasangan
	for i, num := range arr {
		complement := target - num
		if _, ok := numsMap[complement]; ok {
			return []int{numsMap[complement], i}
		}
		numsMap[num] = i
	}

	return []int{}
}

func main() {
	fmt.Println(PairSum([]int{1, 2, 3, 4, 6}, 6)) // [1 3]
	fmt.Println(PairSum([]int{2, 5, 9, 11}, 11))  // [0 2]
	fmt.Println(PairSum([]int{1, 3, 5, 7}, 12))   // [1 3]
	fmt.Println(PairSum([]int{1, 4, 6, 8}, 10))   // [1 2]
	fmt.Println(PairSum([]int{1, 5, 6, 7}, 6))    // []
}
