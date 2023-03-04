package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getPrime(n int) int {
	count := 0
	i := 2
	for {
		if isPrime(i) {
			count++
		}
		if count == n {
			return i
		}
		i++
	}
}

func main() {
	fmt.Println(getPrime(1))
	fmt.Println(getPrime(5))
	fmt.Println(getPrime(8))
	fmt.Println(getPrime(9))
	fmt.Println(getPrime(10))
}
