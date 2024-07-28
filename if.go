package main

import "fmt"

func main() {
	var name = "Lingga"

	if name == "Lingga" {
		fmt.Println("Hello Lingga")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else {
		fmt.Println("Hi, Can I get to know you?")
	}

	if length := len(name); length > 5 {
		fmt.Println("So Long")
	} else {
		fmt.Println("Name is correct")
	}
}
