package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Lingga",
		"address": "Blitar",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Learn Basic Go-Lang"
	book["author"] = "Lingga"
	book["ups"] = "Wrong"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
