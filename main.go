package main

import "fmt"

func getfunc(str string) func(string) string {
	return func(str2 string ) string {
		return str + str2 
	}
}

func main() {
	f1 := getfunc("Hello")
	value := f1("World")
	value2 := f1(" Vikrant")
	fmt.Println(value, value2)
}