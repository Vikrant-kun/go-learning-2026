package main

import "fmt"

type Address struct {
    City string 
    Country string
}

type Developer struct {
    Name string 
    Stack string
    Address
}

func main() {
    dev := Developer{
        Name: "Vikrant",
        Stack: "Go", 
        Address: Address{
            City: "Mumbai",
            Country: "India",
        },
    }
    fmt.Printf("Developer: %s\n", dev.Name)
    fmt.Printf("Stack: %s\n", dev.Stack)
    fmt.Printf("City: %s\n", dev.City)
    fmt.Printf("Country: %s\n", dev.Country)
}
