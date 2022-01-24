package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Loop to", counter)
	}

	slice := []string{"Lingga", "Wahyu", "Rochim", "Budi", "Joko"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for _, value := range slice {
		// fmt.Println("Index", index, "=", value)
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Lingga"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println("Key", key, "=", value)
	}
}
