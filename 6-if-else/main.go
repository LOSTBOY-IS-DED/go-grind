package main

// NOTE : Go does not have a ternary operator like some other languages
// Instead, we use if-else statements for conditional logic

func main() {

	// age := 12

	// if age >= 18 {
	// 	println("You are an adult")
	// } else {
	// 	println("You are a minor")
	// }

	// we can also use like this

	if age := 12; age >= 18 {
		println("You are an adult", age)
	} else {
		println("You are a minor", age)
	}

}
