package main

import (
	"fmt"
)

func dataTypeMap() {
	// inisialisasi basic map
	fmt.Println("\nInitialize data type MAP:")
	var personData = map[string]string{
		"name":     "daffa ganteng mampus",
		"age":      "20",
		"position": "The greatest CTO in Indonesia",
	}
	fmt.Printf("Person data: %v\n", personData)

	fmt.Println("\nMap dengan for range")
	for key, value := range personData {
		fmt.Println(key, "  \t:", value)
	}

	fmt.Println("\nKombinasi slice dan map dan for range")
	var data = []map[string]string{
		{"name": "ayam udin", "age": "setengah abad", "color": "fanta black"},
		{"name": "ayam asep", "age": "setengah abad tambah 2", "color": "yellow black"},
		{"name": "ayam kipli", "age": "setengah juta abad", "color": "red black"},
	}
	for i, item := range data {
		fmt.Printf("ayam %v : %v\n", i+1, item)
	}
}
