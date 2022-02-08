package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Lingga")
	data.PushBack("Wahyu")
	data.PushBack("Rochim")
	data.PushFront("Joko")

	// data.Front().Next().Next().Prev()

	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
