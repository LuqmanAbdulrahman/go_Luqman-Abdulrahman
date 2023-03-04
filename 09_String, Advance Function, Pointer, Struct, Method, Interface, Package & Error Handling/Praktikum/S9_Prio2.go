package main

import "fmt"

func caesar(offset int, input string) string {
	// create an empty string to store the result
	var result string

	// loop over each character in the input string
	for _, char := range input {
		// convert the character to its ASCII code
		code := int(char)

		// apply the Caesar cipher shift
		if code >= 'a' && code <= 'z' {
			code = 'a' + (code-'a'+offset)%26
		} else if code >= 'A' && code <= 'Z' {
			code = 'A' + (code-'A'+offset)%26
		}

		// convert the code back to a character and append it to the result string
		result += string(code)
	}

	// return the result string
	return result
}

func main() {
	fmt.Println(caesar(3, "abc"))                           // def
	fmt.Println(caesar(2, "alta"))                          // cnvc
	fmt.Println(caesar(10, "alterraacademy"))               // kvdobbkkmknowi
	fmt.Println(caesar(1, "abcdefghijklmnopqrstuvwxyz"))    // bcdefghijklmnopqrstuvwxyza
	fmt.Println(caesar(1000, "abcdefghijklmnopqrstuvwxyz")) // mnopqrstuvwxyzabcdefghijkl
}
