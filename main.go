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

//constutor em go (go não tem classes)
// É muito comum utilizar assim, quando cria algo apontando na memoria economiza memoria e evita duplicidade.
func newOrder () *Order {
	return &Order{}
}

func main() {
	a := 1
	b := "b"
	c := 1+1

	//Trabalhando com ponteiros
	d := 10

	// variavel e esta apontando para o endereco de memoria de d, sendo uma referencia de memoria
	e := &d

	// order := Order{
	// 	ID: "123",
	// 	Price: 10.0,
	// 	Quantity: 5,
	// }
	order := newOrder()
	order.ID = "123"
	order.Quantity = 10
	order.Price = 10.0
	order.setPrice(20.0)

	order2 := order
	order2.Price = 5



	fmt.Println(a, b, c, order, order.getTotal(), d, e, *e, order2.Price)
}