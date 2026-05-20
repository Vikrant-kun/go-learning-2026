package main

import "fmt"

func calcBill(pricePerItem float64, quantity int) float64 {
	totalCost := float64(quantity) * pricePerItem
	return totalCost
}

func main() {
	totalCost := calcBill(49.99, 3)
	fmt.Println(totalCost)
}
