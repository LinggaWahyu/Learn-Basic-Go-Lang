package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuu from", a.Name)
}

func main() {
	var lingga Customer
	lingga.Name = "Lingga"
	lingga.Address = "Indonesia"
	lingga.Age = 21

	lingga.sayHi("Joko")
	lingga.sayHuuu()

	// fmt.Println(lingga.Name)
	// fmt.Println(lingga.Address)
	// fmt.Println(lingga.Age)

	// joko := Customer{
	// 	Name:    "Joko",
	// 	Address: "Indonesia",
	// 	Age:     30,
	// }
	// fmt.Println(joko)

	// budi := Customer{"Budi", "Indonesia", 35}
	// fmt.Println(budi)
}
