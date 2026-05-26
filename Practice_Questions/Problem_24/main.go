package main

import "fmt"

func describe(i any) {
	switch v := i.(type) {
	case string:
		fmt.Println("Text:", v)
	case int:
		fmt.Println("Number:", v)
	case bool:
		fmt.Println("Boolean:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	describe("Vikrant")
	describe(22)
	describe(true)
}