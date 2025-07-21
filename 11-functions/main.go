package main

import "fmt"

func main() {
	fmt.Println("Hello we are learning functions")

	simpleFunc()
	ans := add(5, 10)
	fmt.Println("The sum is:", ans)

}

// just a simple function
func simpleFunc() {
	fmt.Println("This is a simple function")
}

// create a add function
func add(a, b int) int { // if the input parameters are of same type, you can write them in a single line if not mentioned separately
	return a + b
}
