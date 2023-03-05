package main

import "fmt"

func generate(numRows int) [][]int {
	result := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0], row[i] = 1, 1
		for j := 1; j < i; j++ {
			row[j] = result[i-1][j-1] + result[i-1][j]
		}
		result[i] = row
	}
	return result
}

func main() {
	fmt.Println(generate(5)) // Output: [[1] [1 1] [1 2 1] [1 3 3 1] [1 4 6 4 1]]
}
