package main

import (
	"fmt"
	"strings"
)

func main() {
	x := "The quick brown fox jumps over the lazy dog"
	fmt.Println("The string used:", x)

	fmt.Println("Length:", len(x))

	//Iterate over each character
	for _, ch := range x {
		fmt.Print(string(ch), ", ")
	}
	fmt.Println()

	fmt.Println("strings.Compare(\"dog\",\"dog\"):", strings.Compare("dog", "dog"))

	fmt.Println("strings.EqualFold(\"dog\",\"Dog\"):", strings.EqualFold("dog", "Dog"))

	x1 := strings.ToLower(x)
	fmt.Println(x1)

	x2 := strings.ToUpper(x)
	fmt.Println(x2)

	x3 := strings.Title(x)
	fmt.Println(x3)
}
