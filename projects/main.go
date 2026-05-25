package main

import "fmt"

type WeightClass string

const (
	Light WeightClass = "Lightweight"
	Medium WeightClass = "Mediumweight"
	Heavy WeightClass = "Heavyweight"
)

type Boxer struct {
	Name        string
    Weight      float64
    WeightClass	WeightClass
    Wins        int
    Losses      int
}

type Gym struct {
	Name    string
	Boxers  []Boxer
}

func (b *Boxer) addWin() {
	b.Wins++
}

func (b *Boxer) addLoss() {
	b.Losses++
}
func (b Boxer) record() string {
	return fmt.Sprintf("%s - %d Wins, %d Losses", b.Name, b.Wins, b.Losses)
}
func (b Boxer) printProfile() {
	fmt.Printf("Name: %s\nWeight: %.2f\nWeight Class: %s\nWins: %d\nLosses: %d\n", b.Name, b.Weight, b.WeightClass, b.Wins, b.Losses)
}

func (g *Gym) addBoxer(b Boxer) {
	g.Boxers = append(g.Boxers, b)
}

func (g Gym) printAll() {
	fmt.Printf("Gym: %s\n", g.Name)
	for _, b := range g.Boxers {
		b.printProfile()
		fmt.Println()
	}
}

func (g Gym) topBoxer() Boxer {
	var top Boxer
	for _, b := range g.Boxers {
		if b.Wins > top.Wins {
			top = b
		}
	}
	return top
}

func (g Gym) byWeightClass(wc WeightClass) []Boxer {
	lightweights := g.byWeightClass(Light)
	fmt.Println("Lightweights:")
	for _, b := range lightweights {
		fmt.Println(b.Name)
	}
	return lightweights
}

type Fighter interface {
    printProfile()
    record() string
}	

func main() {
    gym := Gym{Name: "KD Boxing Academy"}

    gym.addBoxer(Boxer{Name: "Vikrant", Weight: 67.0, WeightClass: Light})
	gym.addBoxer(Boxer{Name: "Rahul", Weight: 72.0, WeightClass: Medium})
	gym.addBoxer(Boxer{Name: "Amit", Weight: 91.0, WeightClass: Heavy})

    gym.Boxers[0].addWin()
    gym.Boxers[0].addWin()
    gym.Boxers[1].addLoss()

    gym.printAll()
    fmt.Println("Top boxer:", gym.topBoxer().Name)
}
