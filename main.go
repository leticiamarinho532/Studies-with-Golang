package main

import "fmt"

// Strut: Forma de estruturar dados
type Order struct {
	ID string
	Price float64
	Quantity int
}

// Metodo da strut Order
func (o Order) getTotal() float64 {
	return o.Price * float64(o.Quantity)
}

func (o *Order) setPrice(price float64) {
	o.Price = price
	fmt.Println("Price dentro do setPrice: ", o.Price)
}

func main() {
	a := 1
	b := "b"
	c := 1+1

	//Trabalhando com ponteiros
	d := 10

	// variavel e esta apontando para o endereco de memoria de d, sendo uma referencia de memoria
	e := &d

	order := Order{
		ID: "123",
		Price: 10.0,
		Quantity: 5,
	}
	order.setPrice(20.0)

	fmt.Println(a, b, c, order, order.getTotal(), d, e, *e)
}