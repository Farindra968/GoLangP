package main

import "fmt"

func main() {
	var nums [4]int
	nums[0]=1

	// prnt full array
	fmt.Println(nums)
	// array indx prin
	fmt.Println(nums[0])
	// Array lengh print
	fmt.Println(len(nums))

	// int
	age := [3]int{20, 24,26}
	name := [3]string{"Ram","Sita","Gita"}

	fmt.Println(name, age)

	// 2d array
	value := [2][2]int{{1,3},{2,4}}
	fmt.Println(value)

}