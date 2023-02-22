package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan%2 == 0 {
		fmt.Println("Bilangan", bilangan, "adalah bilangan genap")
	} else {
		fmt.Println("Bilangan", bilangan, "adalah bilangan ganjil")
	}
}
