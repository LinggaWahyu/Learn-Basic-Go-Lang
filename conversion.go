package main

import "fmt"

func main() {
	var value32 int32 = 257
	var value64 int64 = int64(value32)
	var value8 int8 = int8(value32)

	fmt.Println(value32)
	fmt.Println(value64)
	fmt.Println(value8)

	var name = "Lingga"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}
