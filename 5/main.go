// For Loop
package main

import "fmt"


func main (){
	i:=0

	// simple for loop
	for i<=1000 {
		fmt.Println(i)
		i = i +1
	}

	for i:=1; i<=1000; i++ {
		fmt.Println("Number:",i)
	}

}