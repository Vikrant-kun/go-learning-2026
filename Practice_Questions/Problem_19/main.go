package main

import "fmt"

type Gym struct {
	Name string 
	Location string
}

type Boxer struct {
	Name string 
	Gym Gym 
}

func main() {
	boxer := Boxer{
		Name: "Vikrant",
		Gym: Gym{
			Name: "KD Boxing Academy",
			Location: "Mumbai",
		},
	}
	fmt.Printf("Boxer: %s\n", boxer.Name)
	fmt.Printf("Gym: %s\n", boxer.Gym.Name)
	fmt.Printf("Location: %s\n", boxer.Gym.Location)
}