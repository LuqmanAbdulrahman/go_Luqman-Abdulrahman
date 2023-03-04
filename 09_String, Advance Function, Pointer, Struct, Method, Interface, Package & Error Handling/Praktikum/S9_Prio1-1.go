package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64 // Satuan dalam liter
}

func (c *Car) EstimateDistance() float64 {
	// Menghitung perkiraan jarak yang bisa ditempuh berdasarkan bahan bakar yang terisi
	// dengan asumsi konsumsi bahan bakar 1.5 L/km
	return c.FuelIn / 1.5
}

func main() {
	// Membuat objek Car dengan tipe "sedan" dan bahan bakar terisi 10 liter
	car := Car{Type: "sedan", FuelIn: 10.0}

	// Memanggil method EstimateDistance untuk menghitung perkiraan jarak yang bisa ditempuh
	distance := car.EstimateDistance()

	// Menampilkan hasil perhitungan
	fmt.Printf("Mobil tipe %s dengan bahan bakar terisi %0.2f L dapat menempuh jarak sekitar %0.2f km\n", car.Type, car.FuelIn, distance)
}
