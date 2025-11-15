package main

import "fmt"

func main() {
	// uninitialized slice is deaul nill or null
	// value := []int
	name := []string{"Ram", "Shyam", "Hari", "Laxuman"}
	fmt.Println(cap(name)) // maxium capacity of the slice / like .length in JS
	fmt.Println(name)
}
