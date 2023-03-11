package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
}

func main() {
	// membuat channel untuk menampung hasil pengambilan data produk
	productsCh := make(chan []Product)

	// melakukan panggilan goroutine untuk mengambil data produk
	go func() {
		products, err := getProducts()
		if err != nil {
			panic(err)
		}
		productsCh <- products
	}()

	// menerima hasil pengambilan data produk dari channel
	products := <-productsCh

	// menampilkan data produk
	for _, p := range products {
		fmt.Println("\nProduct Data\n")
		fmt.Println("=====")
		fmt.Printf("ID: %d\n", p.ID)
		fmt.Printf("Title: %s\n", p.Title)
		fmt.Printf("Description: %s\n", p.Description)
		fmt.Printf("Price: $%.2f\n", p.Price)
		fmt.Printf("Category: %s\n", p.Category)
		fmt.Printf("Image: %s\n", p.Image)
		fmt.Println("=====")
	}
}

func getProducts() ([]Product, error) {
	// melakukan panggilan REST API untuk mengambil data produk
	resp, err := http.Get("https://fakestoreapi.com/products")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// memparsing hasil dari response menjadi objek slice produk
	var products []Product
	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
