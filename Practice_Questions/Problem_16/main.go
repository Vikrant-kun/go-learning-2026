package main

import "fmt"

func oddOrEven(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func main() {
	fmt.Println(oddOrEven(7))
	fmt.Println(oddOrEven(12))
	fmt.Println(oddOrEven(33))
	fmt.Println(oddOrEven(100))
}
