package main

import "fmt"

func power(x int, n int) int {
	if n == 0 {
		return 1
	} else if n%2 == 0 {
		temp := power(x, n/2)
		return temp * temp
	} else {
		temp := power(x, (n-1)/2)
		return temp * temp * x
	}
}

func main() {
	fmt.Println(power(2, 3))
	fmt.Println(power(5, 3))
	fmt.Println(power(10, 2))
	fmt.Println(power(2, 5))
	fmt.Println(power(7, 3))

}
