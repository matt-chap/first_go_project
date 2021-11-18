package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	var result string
	switch dow {
	case 1:
		result = "It's Sunday"
		fallthrough
	case 2:
		result = "It's Monday"
		fallthrough
	case 3:
		result = "It's Tuesday"
		fallthrough //since there is no "break" keyword fall through is how to stack cases
	case 4:
		result = "It's Wednesday"
		fallthrough
	case 5:
		result = "It's Thursday"
	case 6:
		result = "It's Friday"
	case 7:
		result = "It's Saturday"
	default:
		result = "Another day"
	}
	fmt.Println(result)
}
