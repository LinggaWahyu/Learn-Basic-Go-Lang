package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Lingga Wahyu", "Lingga"))
	fmt.Println(strings.Contains("Lingga Wahyu", "Joko"))

	fmt.Println(strings.Split("Lingga Wahyu Rochim", " "))

	fmt.Println(strings.ToLower("Lingga Wahyu Rochim"))
	fmt.Println(strings.ToUpper("Lingga Wahyu Rochim"))
	fmt.Println(strings.ToTitle("lingga wahyu rochim"))

	fmt.Println(strings.Trim("        Lingga Wahyu        ", " "))

	fmt.Println(strings.ReplaceAll("Joko Joko Joko Joko", "Joko", "Lingga"))

}
