package main

import (
	"fmt"
)

func main() {
	var word string
	fmt.Print("Masukkan kata: ")
	fmt.Scanln(&word)

	if isPalindrome(word) {
		fmt.Printf("%s adalah palindrome\n", word)
	} else {
		fmt.Printf("%s bukan palindrome\n", word)
	}
}

func isPalindrome(word string) bool {
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}
	return true
}
