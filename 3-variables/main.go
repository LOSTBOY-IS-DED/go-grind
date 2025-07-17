package main

// note : Its necessary to use every variable declared in the code, otherwise it will throw an error

func main() {

	var name string = "Subh"   // this is how to declare a variable though there is a shorter way
	var lastName = "Chaudhury" // type inference
	println(name)
	println(lastName)
	var isAwesome = true
	println(isAwesome) // prints the type of the variable

	// shorthand syntax
	fullName := "Subh Chaudhury" // we use this whenever we want to declare and assign a variable in one line
	println(fullName)

	var girlFriend string
	girlFriend = "Neha"

	println(girlFriend) // prints the value of the variable

	var price float64 = 10.99
	println(price)
}
