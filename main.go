package main

import "fmt"

type Animal interface {
	Says() string
	NumberofLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorila struct {
	Name          string
	Color         string
	NumberofTeeth int
}

func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shephered",
	}

	gorilla := Gorila{
		Name:          "Gorilla",
		Color:         "Black",
		NumberofTeeth: 30,
	}

	PrintInfo(&dog)
	PrintInfo(&gorilla)
}

func PrintInfo(a Animal) {
	fmt.Println("This animal say", a.Says(), "and has", a.NumberofLegs(), "legs")
}

func (d *Dog) Says() string {
	return "Woof"
}
func (d *Dog) NumberofLegs() int {
	return 4
}

func (d *Gorila) Says() string {
	return "Uggh"
}
func (d *Gorila) NumberofLegs() int {
	return 2
}
