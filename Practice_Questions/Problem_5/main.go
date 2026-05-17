package main

import "fmt"

type lineup	[]string

func (l lineup) print() {
	fmt.Println("First:", l[0])
	fmt.Println("Last:", l[len(l)-1])
}


func main() {
	names := lineup{"Mango", "Apple", "Banana", "Orange", "Grape"}
	names.print()
}