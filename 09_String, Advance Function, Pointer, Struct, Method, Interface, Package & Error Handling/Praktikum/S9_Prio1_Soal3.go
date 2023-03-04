package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var common string

	// cari substring yang sama pada kedua string
	for i := 0; i < len(a); i++ {
		for j := i + 1; j <= len(a); j++ {
			if strings.Contains(b, a[i:j]) && len(a[i:j]) > len(common) {
				common = a[i:j]
			}
		}
	}

	return common
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))
}
