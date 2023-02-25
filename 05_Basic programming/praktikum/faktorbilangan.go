package main

import "fmt"

func main() {
	var number int
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&number)

	fmt.Printf("Faktor dari %d adalah: ", number)

	// loop melalui semua bilangan dari 1 hingga number
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
}
