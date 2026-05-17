package main

import "fmt"

type digits []int

func (d digits) countFives() {
    count := 0
    for _, num := range d {
        if num == 5 {
            count++
        }
    }
    fmt.Println("There are", count, "fives in the list.")
}

func main() {
    num := digits{5, 3, 5, 7, 5, 2, 5, 8}
    num.countFives()
}