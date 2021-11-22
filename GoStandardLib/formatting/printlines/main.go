package main

import "fmt"

func main() {
	const a, b, c = 5, 6, 11

	length, err := fmt.Println("Adding", a, "and", b, "you get", c)

	fmt.Println("The previous statement had", length, "characters and", err, "errors.")
}
