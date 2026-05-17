package main

import "fmt"

type letters []string

func (l letters) reversePrint() {
	for i := len(l)-1; i>=0; i-- {
		fmt.Println(l[i])
	}
}

func main() {
	list := letters{"A", "B", "C", "D"}
	list.reversePrint()
}