package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"

	fmt.Println(colors)
	fmt.Println(colors[0])

	var numbers = [5]int{0, 1, 2, 3, 4}
	fmt.Println(numbers)
	fmt.Println("Length of array", len(numbers))
}
