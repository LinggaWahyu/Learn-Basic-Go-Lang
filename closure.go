package main

import "fmt"

func main() {
	counter := 0
	name := "Lingga"

	increment := func() {
		fmt.Println("Increment")
		counter++
		name := "Joko"

		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
