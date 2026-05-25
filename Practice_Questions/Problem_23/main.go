package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct {}
type Bike struct {}
type Boat struct {}

func (c Car) move() {
	fmt.Println("driving")
}
func (b Bike) move() {
	fmt.Println("riding")
}
func (b Boat) move() {
	fmt.Println("sailing")
}

func identify(v Vehicle) {
	switch v.(type) {
	case Car:
		fmt.Println("This is a car - driving")
	case Bike:
		fmt.Println("This is a bike - riding")
	case Boat:
		fmt.Println("This is a boat - sailing")
	}
}

func main() {
    identify(Car{})
    identify(Bike{})
    identify(Boat{})
}