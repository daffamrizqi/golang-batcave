package main

import (
	"fmt"
)

func main() {

	// usual for loop
	fmt.Println("Usual for loop")
	for i := 0; i <= 10; i++ {
		fmt.Printf("Counter: %v\n", i)
	}

	// for dengan argumen hanya kondisi
	fmt.Println("\nfor dengan argumen hanya kondisi")
	i := 0
	for i < 10 {
		fmt.Printf("Counter: %v\n", i)
		i++
	}
}
