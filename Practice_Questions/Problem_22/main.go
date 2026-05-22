package main

import "fmt"

type Animal interface {
    speak() string 
}

type Dog struct {}
type Cat struct {}
type Cow struct {}  

func (d Dog) speak() string {
    return "Woof!"
}

func (c Cat) speak() string {
    return "Meow!"
}

func (c Cow) speak() string {
    return "Moo!"
}

func makeNoise(a Animal) {
    fmt.Println(a.speak())
}   

func main() {
    makeNoise(Dog{})
    makeNoise(Cat{})
    makeNoise(Cow{})
}