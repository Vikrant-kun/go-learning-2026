package main

import "fmt"

func totalFightTime(rounds int, minutesPerRound int) int {
	totalMinutes := rounds * minutesPerRound
	return totalMinutes
}

func main() {
	result := totalFightTime(12, 3)
	fmt.Println("Total fight time:", result, "minutes")
}
