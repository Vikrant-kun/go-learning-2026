package main

import "fmt"

type scores []int

func (s scores) highest() int {
    highest := 0
    for _, max := range s {
        if max > highest {
            highest = max
        }
    }
    fmt.Println("Highest score:", highest)
    return highest
}

func main() {
    s := scores{85, 92, 78, 90, 88}
    fmt.Println("The highest score is:", s.highest())
}