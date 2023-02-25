package main

import "fmt"

func main() {
	var sisiAtas, sisiBawah, tinggi float64

	fmt.Print("Masukkan panjang sisi atas: ")
	fmt.Scanln(&sisiAtas)

	fmt.Print("Masukkan panjang sisi bawah: ")
	fmt.Scanln(&sisiBawah)

	fmt.Print("Masukkan tinggi: ")
	fmt.Scanln(&tinggi)

	luas := (sisiAtas + sisiBawah) * tinggi / 2

	fmt.Printf("Luas trapesium adalah: %.2f", luas)
}
