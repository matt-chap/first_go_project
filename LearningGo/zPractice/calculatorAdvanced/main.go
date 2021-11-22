package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

//TODO: Add 2 values and let the user chose an operation
func main() {

	float1 := GetUserInput("Enter Value 1: ")
	float2 := GetUserInput("Enter Value 2: ")
	DoTheMath("Select an operation (+ - * /):", float1, float2)

	//If you want to build to new OS simply use which one you want (darwing is apple)
	//GOOS="window" go buid
	//GOOS="darwin" go buid
	//GOOS="linux" go buid

}

func GetUserInput(label string) float64 {
	fmt.Print(label)
	input1, _ := reader.ReadString('\n')
	userInput, err1 := strconv.ParseFloat(strings.TrimSpace(input1), 64)

	if err1 != nil {
		panic("You broke the calculator by adding a string.")
	}

	return userInput
}

func DoTheMath(label string, float1 float64, float2 float64) {
	fmt.Print(label)
	input1, _ := reader.ReadString('\n')
	userInput := strings.TrimSpace(input1)

	switch userInput {
	case "+":
		fmt.Printf("The sum of %v and %v rounded to the nearest hundredth place equals %v\n", float1, float2, math.Round((float1+float2)*100)/100)
	case "-":
		fmt.Printf("The differance of %v and %v rounded to the nearest hundredth place equals %v\n", float1, float2, math.Round((float1-float2)*100)/100)
	case "*":
		fmt.Printf("The product of %v and %v rounded to the nearest hundredth place equals %v\n", float1, float2, math.Round((float1*float2)*100)/100)
	case "\\":
		fallthrough
	case "/":
		fmt.Printf("The quotient of %v and %v rounded to the nearest hundredth place equals %v\n", float1, float2, math.Round((float1/float2)*100)/100)
	default:
		panic("Not able to recognize the operation")
	}
}
