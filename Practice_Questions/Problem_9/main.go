package main

import (
	"fmt"
)

type goals []int

func (g goals) total() {
	sum := 0
	for _, value := range g {
		sum += value
	}
	fmt.Println("Total is : ", sum)
}

func (g goals) highest() {
	num := 0
	for _, max := range g {
		if max > num {
			num = max
		}
	}
	fmt.Println("Highest is :", num)
}

func main () {
	g := goals{5, 10, 15, 20, 25}
	g.total()
	g.highest()
}