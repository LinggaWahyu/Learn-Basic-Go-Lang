package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Change"
	// fmt.Println(slice1)

	// slice1[0] = "May Again"
	// fmt.Println(months)

	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "New Month")
	fmt.Println(slice3)
	slice3[1] = "Not December"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Lingga"
	newSlice[1] = "Wahyu"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	isArray := [5]int{1, 2, 3, 4, 5}
	isSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(isArray)
	fmt.Println(isSlice)
}
