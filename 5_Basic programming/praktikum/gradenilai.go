package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("Masukkan nilai Anda: ")
	fmt.Scan(&nilai)

	if nilai < 0 || nilai > 100 {
		fmt.Println("Nilai Invalid")
	} else if nilai >= 80 {
		fmt.Println("Grade: A")
	} else if nilai >= 65 {
		fmt.Println("Grade: B")
	} else if nilai >= 50 {
		fmt.Println("Grade: C")
	} else if nilai >= 35 {
		fmt.Println("Grade: D")
	} else {
		fmt.Println("Grade: E")
	}
}
