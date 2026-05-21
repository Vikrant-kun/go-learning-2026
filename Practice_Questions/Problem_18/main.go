package main

import "fmt"

type Boxer struct {
	name   string
	weight float64
	wins   int
}

func (b Boxer) introduce() {
	fmt.Println("I am", b.name, b.weight, "kg,", b.wins, "wins")
}

func main() {
	b := Boxer{
		name:   "Vikrant",
		weight: 67.0,
		wins:   10,
	}
	b.introduce()
}
