package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (c *Car) EstimateDistance() float64 {
	return c.FuelIn * 1.5
}

func main() {
	myCar := Car{Type: "SUV", FuelIn: 20}
	estimatedDistance := myCar.EstimateDistance()
	fmt.Printf("Estimated distance for %s with %.2f L fuel: %.2f mill\n", myCar.Type, myCar.FuelIn, estimatedDistance)
}
