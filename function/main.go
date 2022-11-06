package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// simple example function
	var fullName = []string{"john", "wick"}
	printMessage("hello", fullName)

	// function with return value
	fmt.Println("\nfunction example with return value:")
	rand.Seed(time.Now().Unix())
	for i := 1; i <= 3; i++ {
		fmt.Printf("Try number %v: %v\n", i+1, randomWithRange(1, 10))
	}

	// stoping function with return
	fmt.Println("\nStopping function with return keyword:")
	for i := -1; i <= 3; i++ {
		divideNumber(randomWithRange(1, 10), i)
		fmt.Print("\n")
	}

	// function returning multiple values
	var input float64 = 7
	circleArea, circelCifcumference := circleCalculation(input)
	fmt.Printf("\ncicle area with %v diameter: %v\n", input, circleArea)
	fmt.Printf("cicle circumference with %v diameter: %v\n", input, circelCifcumference)

	// variadic function
	fmt.Println("\nVariadic function:")
	var numbers = []int{2, 3, 4, 8, 3, 4, 5, 6, 2}
	avg, totalData, totalScore := calculateAvg(numbers...)
	fmt.Printf("Total Data: %v\nTotal Score: %v\nAverage: %v\n", totalData, totalScore, avg)

	// function closure:
	fmt.Println("\nContoh function closure:")
	getMinMax := func(nums []int) (int, int) {
		var min, max int
		for i, x := range nums {
			switch {
			case i == 0:
				max, min = x, x
			case x > max:
				max = x
			case x < min:
				min = x
			}
		}
		return min, max
	}
	min, max := getMinMax(numbers)
	fmt.Printf("data: %v\nmin: %v\nmax: %v\n", numbers, min, max)

	// immidiate-invoked function expression
	newNumbers := func(min int) []int {
		var newArrar []int
		for _, num := range numbers {
			if num < min {
				continue
			}
			newArrar = append(newArrar, num)
		}
		return newArrar
	}(3)
	fmt.Println("\nImidiate-invoked-function expression:")
	fmt.Println("Original numbers :", numbers)
	fmt.Println("Filtered numbers :", newNumbers)
}

func printMessage(message string, fullName []string) {
	nameString := strings.Join(fullName, " ")
	fmt.Println(message, nameString)
}

// function returning value
func randomWithRange(min, max int) int {
	value := rand.Int()%(max-min+1) + min
	return value
}

// stopping function with return keyword
func divideNumber(num1, num2 int) {
	if num2 == 0 {
		fmt.Print("Cannot divide by 0")
		return
	}
	result := num1 / num2
	fmt.Printf("%v / %v = %v", num1, num2, result)
}

// function returning multiple value
func circleCalculation(d float64) (float64, float64) {
	area := math.Pi * math.Pow(d/2, 2)
	circumference := math.Pi * d
	return area, circumference
}

// variadic function
func calculateAvg(numbers ...int) (float64, int, int) {
	var totalData int
	for _, num := range numbers {
		totalData += num
	}

	avg := totalData / len(numbers)
	totalScore := len(numbers)
	return float64(avg), totalData, totalScore
}
