package main

import "fmt"

func main() {
	x := 10

	//Decimal
	fmt.Printf("%d\n", x)

	//hexadecimal
	fmt.Printf("%x\n", x)

	//bool
	fmt.Printf("%t\n", x > 10)

	//float
	fmt.Printf("%f\n", 10.23)
	fmt.Printf("%e\n", 10.23)
	fmt.Printf("%010.2f\n", 10.12345)

	//explicit arguments
	fmt.Printf("%[2]d %[1]d\n", 22, 33)

	//print struct with properties
	c := circle{
		radius: 10,
		border: 2,
	}
	fmt.Printf("%+v\n", c)
}

type circle struct {
	radius int
	border int
}
