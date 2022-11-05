package main

import (
	"fmt"
)

func main() {
	fName := "Daffa ganteng"
	lName := "Rizqi jago ngoding"

	age := 20

	fmt.Printf("Hello my name is %s %s and I am %v years old\n", fName, lName, age)

	var randomNum uint8 = 255
	fmt.Printf("ini adalah random uint8 = %v\n", randomNum)

	// variable temporary
	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}
}
