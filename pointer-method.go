package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	lingga := Man{"Lingga"}
	lingga.Married()
	fmt.Println(lingga.Name)
}
