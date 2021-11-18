package main

import (
	"fmt"
)

func main() {
	var colors = []string{"Red", "Blue", "Green"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	// for can be a while true
	value := 1
	for value < 10 {
		fmt.Println(value)
		value++
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println(sum)
		if sum > 200 {
			goto theEnd
		}
	}

theEnd:
	fmt.Println("End of Program")
}
