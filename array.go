package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Lingga"
	names[1] = "Wahyu"
	names[2] = "Rochim"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int {
		80,
		95,
		90,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var again [10]string
	fmt.Println(len(again))

}
