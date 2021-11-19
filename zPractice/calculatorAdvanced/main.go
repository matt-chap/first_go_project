package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//TODO: Add 2 values and let the user chose an operation
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Value 1: ")
	input1, _ := reader.ReadString('\n')
	float1, err1 := strconv.ParseFloat(strings.TrimSpace(input1), 64)

	if err1 != nil {
		panic("You broke the calculator by adding a string.")
	}

	fmt.Print("Enter Value 2: ")
	input2, _ := reader.ReadString('\n')
	float2, err2 := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	if err2 != nil {
		panic("You broke the calculator by adding a string.")
	}

	fmt.Printf("The sum of %v and %v rouanded to the nearest hundredth place equals %v\n", float1, float2, math.Round((float1+float2)*100)/100)

	//If you want to build to new OS simply use which one you want (darwing is apple)
	//GOOS="window" go buid
	//GOOS="darwin" go buid
	//GOOS="linux" go buid

}
