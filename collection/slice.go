package main

import (
	"fmt"
)

func slice() {
	// initialize slice
	fmt.Println("\ninitialize slice and using range method iterate:")
	var nums = []int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Printf("No. %v\n", num)
	}

	// 2 indexing access slice
	fmt.Println("\nslice technique: 2 index")
	newNums := nums[0:2]
	for _, num := range newNums {
		fmt.Printf("newNums: %v\n", num)
	}
}
