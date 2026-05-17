package main

import "fmt"

type items []string

func (d items) evenIndex() {
	for i, list := range d {
		if i % 2 == 0 {
			fmt.Println(i, list)
			}
	}
}

func main() {
	list := items{"A", "B", "C", "D", "E", "F"}
	list.evenIndex()
}