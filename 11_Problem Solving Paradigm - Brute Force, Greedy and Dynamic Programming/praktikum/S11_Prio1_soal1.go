package main

import "fmt"

func getBinaryArray(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ans[i] = getBinary(i)
	}
	return ans
}

func getBinary(num int) int {
	if num == 0 {
		return 0
	}
	binary := 0
	base := 1
	for num > 0 {
		rem := num % 2
		binary += rem * base
		num = num / 2
		base = base * 10
	}
	return binary
}

func main() {
	fmt.Println(getBinaryArray(2)) // Output: [0 1 10]
	fmt.Println(getBinaryArray(3)) // Output: [0 1 10 11]
}
