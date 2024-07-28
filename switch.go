package main

import "fmt"

func main() {
	name := "Lingga"

	switch name {
	case "Lingga":
		fmt.Println("Hello Lingga")
		fmt.Println("Hello Lingga")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hi, Can I get to know you?")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("So Long")
	// case false:
	// 	fmt.Println("Name is correct")
	// }

	length := len(name)

	switch {
	case length > 10:
		fmt.Println("So Long")
	case length > 5:
		fmt.Println("Quite Long")
	default:
		fmt.Println("Name is correct")
	}
}
