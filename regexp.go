package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("l([a-z])n")

	fmt.Println(regex.MatchString("lin"))
	fmt.Println(regex.MatchString("lan"))
	fmt.Println(regex.MatchString("lUn"))

	search := regex.FindAllString("lin lan len lun lon lum", -1)
	fmt.Println(search)
}
