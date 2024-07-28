package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error with message", message)
	}
	fmt.Println("Application Done")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APPLICATION ERROR")
	}
	fmt.Println("Application Running")
}

func main() {
	runApp(false)
}
