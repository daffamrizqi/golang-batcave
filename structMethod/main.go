package main

import (
	"fmt"
)

// deklarasi struct
type engineer struct {
	name     string
	position string
	age      int
}

func main() {
	// penerapan struct engineer
	var daffa engineer
	daffa.name = "Daffa Muhammad Rizqi"
	daffa.position = "CTO"
	daffa.age = 20
	fmt.Printf("Name: %v\n", daffa.name)
	fmt.Printf("Position: %v\n", daffa.position)

	// Inisialisasi Object Struct
	var michael = engineer{"Ethan", "frontend", 22}
	fmt.Println("ini michael", michael)

}
