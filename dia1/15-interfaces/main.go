package main

import "fmt"

// An interface type is defined as a set of method signatures.
//   A value of interface type can hold any value that implements those methods.

// Animal is interface for elements of Zoo
type Animal interface {
	Shout() string
}

// Cow implementa o que a interface animal especifica
type Cow struct{}

// Shout return sound by animal
func (c *Cow) Shout() string {
	return "Muuu"
}

// Dog implementa o que a interface animal especifica
type Dog struct{}

// Shout return sound by animal
func (d *Dog) Shout() string {
	return "Rau rau!"
}

func main() {
	var animal Animal

	animal = &Cow{}
	callMyAnimal(animal)

	animal = &Dog{}
	callMyAnimal(animal)
}

func callMyAnimal(a Animal) {
	fmt.Println("My animal says:", a.Shout())
}
