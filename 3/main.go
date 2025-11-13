package main

import "fmt"

// Variable and data types

func main() {
	var name string = "Farindra"
	var age int = 25
	var height = 5.6 // type inferred [auto type detect] 

	// shorthand method
	price := 30.5  // auto declare variable & its type

	var add string
	add = "56"
	var country = "Nepal"
	const Pi = 5.14 // constant
	const Greeting = "Namaste" // constant
	fmt.Println(name, age, height, add, price, country, Pi, Greeting)
}