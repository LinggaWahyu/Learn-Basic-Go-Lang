package main

import "fmt"

func main() {

	var finalScore = 80
	var absence = 80

	var passExam = finalScore >= 80
	var passAbsence = absence >= 80

	var pass = passExam && passAbsence
	fmt.Println(pass)

	fmt.Println(finalScore >= 80 && absence >= 80)

}
