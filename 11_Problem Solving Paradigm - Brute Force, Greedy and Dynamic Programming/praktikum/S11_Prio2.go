package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	n := len(jumps)
	dp := make([]int, n)
	dp[0] = 0
	dp[1] = int(math.Abs(float64(jumps[1] - jumps[0])))

	for i := 2; i < n; i++ {
		dp[i] = int(math.Min(float64(dp[i-1]+int(math.Abs(float64(jumps[i]-jumps[i-1])))),
			float64(dp[i-2]+int(math.Abs(float64(jumps[i]-jumps[i-2]))))))
	}
	return dp[n-1]
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}
