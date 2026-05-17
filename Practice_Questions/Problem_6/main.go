package main

import "fmt"

type prices []float64

func (d prices) print() { 
	for _, multi := range d {
		fmt.Println(multi*2)
		}
	}


func main() {
	multi := prices{10.5, 20.0, 35.5, 50.0}
	multi.print()
}