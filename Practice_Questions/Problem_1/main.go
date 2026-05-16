package main

import "fmt"

type number []int
func (n number) sum() int {
    total := 0
    for _, value := range n {
        total += value
    }
    fmt.Println("Total is", total)
    return total
}

func main() {
    n := number{10, 20, 30, 40}
    n.sum()
}

