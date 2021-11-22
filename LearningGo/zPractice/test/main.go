package main

import "fmt"

func main() {
	var out int
	for j := 0; j < 20; j++ {
		fmt.Println("out", out)
		fmt.Println("j*j", j*j)
		out = j*j + out
		if out > 12 {
			goto theEnd
		}
	}
theEnd:
	fmt.Println(out)
}
