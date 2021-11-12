package main

import "fmt"

const constString = "This string is a const"

func main() {
	//printing the const and since it is lowercase it is protected
	fmt.Println(constString)
	fmt.Printf("The variables type is %T\n", constString)
	fmt.Println()

	//implicit typing (explicit would have "string" key word after the aString)
	var aString = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("The variables type is %T\n", aString)
	fmt.Println()

	//explicit typing (there is diference between int on different OS's so simply call out which one to avoid any confusion)
	var anInt int32 = 42
	fmt.Println(anInt)
	fmt.Printf("The variables type is %T\n", anInt)
	fmt.Println()

	//short hand (can only be used within a function)
	newString := "This is a new string"
	fmt.Println(newString)
	fmt.Printf("The variables type is %T\n", newString)
	fmt.Println()
}
