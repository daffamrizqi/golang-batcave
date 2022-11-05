package main

import (
	"fmt"
)

func main() {
	slice()
	dataTypeMap()

	// initialize array
	var names [3]string
	names[0] = "dodi"
	names[1] = "udin petot"
	names[2] = "ujang pekok"

	// initialize array dengan keyword make
	fmt.Println("\nArray dengan keyword make")
	var brands = make([]string, 3)
	brands[0] = "nike"
	brands[1] = "adidas"
	for i, brand := range brands {
		if brand != "" {
			fmt.Printf("brand %v: %v\n", i+1, brand)
		}
	}

	// inisialisasi array dengan nilai awal array
	fmt.Println("\ninisialisasi array dengan nilai awal array:")
	var nums = [4]int{1, 2, 3}
	for _, num := range nums {
		if num != 0 {
			fmt.Printf("No. %v\n", num)
		}
	}

	// inisialisasi nilai awal array tanpa jumlah elemen
	fmt.Println("\nInisialisai nilai awal array tanpa jumlah elemen:")
	var coffees = [...]string{
		"arabica",
		"robusta",
		"liberica",
	}
	for i, coffee := range coffees {
		fmt.Printf("%v. %v\n", i+1, coffee)
	}

	//
}
