package main

import (
	"fmt"
)

func main() {
	doSomething()

	sum := addValues(5, 8)
	fmt.Println(sum)

	multiValues := addAllValues(1, 2, 3, 55, 53)
	fmt.Println("multivalues:", multiValues)

	multiTotal, multiLen := returnMultiValues(1, 2, 3, 55, 53)
	fmt.Printf("total: %v\nlength %v", multiTotal, multiLen)
}

func doSomething() {
	fmt.Println("Doing Something")
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func returnMultiValues(values ...int) (int, int) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, len(values)
}
