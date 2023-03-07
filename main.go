package main

import "fmt"

// Strut: Forma de estruturar dados
type Order struct {
	ID string
	Price float64
	Quantity int
}

func main() {
	a := 1
	b := "b"
	c := 1+1

	order := Order{
		ID: "123",
		Price: 10.0,
		Quantity: 5,
	}

	fmt.Println(a, b, c, order)
}