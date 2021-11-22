package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 15.3}
	//String format of a struct
	fmt.Println("Dog:", poodle)

	//formated stuct
	fmt.Printf("Dog: %+v\n", poodle)
	fmt.Printf("Dog Breed: %v\n", poodle.Breed)

	poodle.Weight = 16
	fmt.Printf("Dog New Weight: %v\n", poodle.Weight)
}

// uppercase means its exported for the struct
// upercase on the properties also means its exported
type Dog struct {
	Breed  string
	Weight float32
}
