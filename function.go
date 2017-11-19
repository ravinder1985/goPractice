package main

import "fmt"

func compare(T int) {
	switch T {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("NONE")
	}
}

func rev(number int) int {
	div := 10
	reminder := 1
	newNumber := 0
	for reminder > 0 {
		reminder = number % div
		number = number - reminder
		number = number / 10
		if reminder > 0 {
			newNumber = newNumber * 10
			newNumber = newNumber + reminder
		}
	}
	return newNumber
}
