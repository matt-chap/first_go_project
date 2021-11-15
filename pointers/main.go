package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var pointer1 = &anInt
	fmt.Println("Value of pointer1:", *pointer1)

	aFloat := 42.3
	pointer2 := &aFloat
	fmt.Println("Value of pointer2:", *pointer2)

	aVal := 42.3
	pointer3 := &aVal
	fmt.Println("Value of pointer3:", *pointer3)

	*pointer3 = *pointer3 / 2
	fmt.Println("Value of pointer3:", *pointer3)
	fmt.Println("Value of aVal:", aVal)
}
