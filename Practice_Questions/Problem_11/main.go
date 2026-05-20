package main

import "fmt"

func canVote(age int) bool {
	if age>=18 {
		return true
	} else {
		return false 
	}
}

func main() {
	result := canVote(12)
	fmt.Println(result)
}