package main

import (
	"fmt"
	"time"
)

func main() {
	// simple witch
	i := 1
	switch i {
	case 1:
		fmt.Println("Value ", i)
	case 2:
		fmt.Println("value", i)
	case 3:
		fmt.Println("value", i)
	}

	// multi condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Friday:
		fmt.Println("Todays is Weekend")
	default:
		fmt.Println("Today is workday ")
	}
}
