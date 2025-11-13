// For Loop
package main

import "fmt"

func main() {
	i := 0

	// simple for loop
	for i <= 1000 {
		fmt.Println(i)
		i = i + 1
	}

	// classic or loop
	for i := 1; i <= 1000; i++ {
		fmt.Println("Number:", i)
	}

	// for loop rang
	for i := range 100 {
		fmt.Println(i) // range help to print specific range starting form 0  // e.g range 100 - 0-99
	}

	// infinite for loop
	for {
		fmt.Println("Namasta")
	}

}
