package main

import "fmt"

func repeatPrint(message string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(message)
	}
}

func main() {
	repeatPrint("Go is awesome!", 7)
}
