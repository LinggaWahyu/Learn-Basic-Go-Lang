package main

import (
	"Learn-Basic-Golang/helper"
	"fmt"
)

func main() {
	helper.SayHello("Lingga")
	// helper.sayGoodBye("Lingga") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}
