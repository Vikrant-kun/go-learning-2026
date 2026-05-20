package main

import "fmt"

func celsiusToFahrenheit(celcius float64) float64 {
	fahrenheit := (celcius * 9 / 5) + 32
	return fahrenheit
}

func main() {
	convert := celsiusToFahrenheit(37)
	fmt.Println(convert)
}
