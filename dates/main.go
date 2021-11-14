package main

import (
	"fmt"
	"time"
)

func main() {

	n := time.Now()
	fmt.Println("Date now:", n)

	t := time.Date(1992, time.December, 25, 4, 0, 0, 0, time.UTC)
	fmt.Println("Date created:", t)
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("Jan 01 2006 3:04 PM")) // Go uses weird consts for date formating

	parsedTime, _ := time.Parse(time.ANSIC, "Fri Dec 25 04:00:00 1992")
	fmt.Printf("The type of the parsedTime object is %T\n", parsedTime)
}
