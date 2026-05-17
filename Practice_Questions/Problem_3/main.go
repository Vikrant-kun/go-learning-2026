package main

import "fmt"

type numList []int

func (n numList) bigones() {
    for _, num := range n {
        if num > 50 {
            fmt.Println(num)
        }
    }
}

func main() {
    numbers := numList{30, 67, 45, 89, 12, 55, 78}
    numbers.bigones()
}