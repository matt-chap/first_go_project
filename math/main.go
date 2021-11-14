package main

import (
	"fmt"
	"math"
)

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3

	fmt.Println("Interger sum:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.32
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("Float rounded:", floatSum)

	floatSum = 32.456235256722
	fmt.Printf("Float rounded: %.2f\n", floatSum)
}
