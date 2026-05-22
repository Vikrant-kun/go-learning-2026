package main

import "fmt"

type Shape interface {
    area() float64 
}

type Circle struct {
    radius float64
}

type Rectangle struct {
    width float64
    height float64
}

func (c Circle) area() float64 {
    return 3.14 * c.radius * c.radius
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

func printArea(s Shape) {
    fmt.Printf("Area: %.2f\n", s.area())
}

func main() {
    c := Circle{radius: 5}
    r := Rectangle{width: 4, height: 6}

    printArea(c)  // passes Circle as Shape
    printArea(r)  // passes Rectangle as Shape
}
