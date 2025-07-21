package main

import "fmt"

// dynamic sized arrays
// most used data structure in go
// slices are more flexible than arrays
// useful methods like append, copy, etc.

func main() {
	// NOTE ; uninitialized slices are nil
	var nums []int

	fmt.Println(nums)      // prints []
	fmt.Println(len(nums)) // prints 0

	var num2 = make([]int, 2, 5) // creates an empty slice with length 2 and capacity 5
	fmt.Println(len(num2))       // prints 2
	// capacity is the maximum size of the slice
	fmt.Println(cap(num2)) // prints [0 0]
	num2 = append(num2, 1) // appending value to the slice at the end --> capacity is increased itself when needed
	fmt.Println(num2)      // prints [0 0 1]
}
