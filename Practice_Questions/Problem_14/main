package main

import "fmt"

func monthlySavings(income float64, expenses float64) float64 {
	savings := income - expenses
	if expenses > income {
		return 0
	}
	return savings
}

func main() {
	money := monthlySavings(50000, 35000)
	fmt.Println(money)
}
