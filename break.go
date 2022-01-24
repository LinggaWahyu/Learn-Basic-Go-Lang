package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Loop to", i)
		if i == 5 {
			break
		}
	}
}
