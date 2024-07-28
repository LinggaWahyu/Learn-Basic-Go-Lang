package main

import "fmt"

func main() {
	type numberID string
	type Married bool

	var numberIdLingga numberID = "12345667890"
	var marriedStatus Married = true
	fmt.Println(numberIdLingga)
	fmt.Println(marriedStatus)
}
