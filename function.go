package main

import "fmt"
import "os"
import "strconv"

func add(x , y int) int {
	return x + y
}

func rev(number int) int{
	div := 10
	reminder := 1
	newNumber := 0
	for reminder > 0 {
		reminder = number % div
		number = number - reminder
		number = number/10
		if reminder > 0  {
			newNumber = newNumber * 10
			newNumber = newNumber + reminder
		}
	}
	return newNumber
}

func main(){
	a, _ := strconv.Atoi(os.Args[1])
	fmt.Println(rev(int(a)))
}
