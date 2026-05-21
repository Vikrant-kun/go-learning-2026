package main

import "fmt"

func monthlyFee(sessionsPerWeek int, feePerSession float64) float64 {
	monthlyCost := float64(sessionsPerWeek) * 4 * feePerSession
	return monthlyCost
}

func main() {
	monthlyCost := monthlyFee(6, 500.0)
	fmt.Println("Monthy boxing fee: ", monthlyCost)
}
