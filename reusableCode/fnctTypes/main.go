package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 15.3, "Woof!"}
	//String format of a struct
	fmt.Println("Dog:", poodle)

	//formated stuct
	fmt.Printf("Dog: %+v\n", poodle)
	fmt.Printf("Dog Breed: %v\n", poodle.Breed)

	poodle.Weight = 16
	fmt.Printf("Dog New Weight: %v\n", poodle.Weight)

	poodle.Speak()
	poodle.Sound = "Bark"
	poodle.Speak()

	refPoodle := &poodle
	poodle.SpeakThreeTimes()

	refPoodle.PointerDog()
	refPoodle.PointerDog()
}

// uppercase means its exported for the struct
// upercase on the properties also means its exported
type Dog struct {
	Breed  string
	Weight float32
	Sound  string
}

func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

//This is pass by value and not by reference sooooo
//If you call this again from the code it only barks 3 times
func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}

//there is no overloading so it has to have a new name
func (d *Dog) PointerDog() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}
