package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Blue", "Green"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)

	// they shortened colors = append(colors[1:len(colors)])
	colors = append(colors[1:])
	fmt.Println(colors)

	//capacity is optional
	numbers := make([]int, 5, 5)
	numbers[0] = 23
	numbers[1] = 734
	numbers[2] = 233
	numbers[3] = 34
	numbers[4] = 5
	fmt.Println(numbers)

	//sorting
	sort.Ints(numbers)
	fmt.Println(numbers)

}
