package main

// NOTE : you can also declare variables outside the main function, but they must be used in the main function or it will throw an error

func main() {
	const name = "Subh"
	const lastName = "Chaudhury"

	// name = "Subh Chaudhury" // This will throw an error because constants cannot be reassigned
	println(name)
	println(lastName)

	const (
		port = 8080
		host = "localhost"
	)

	println(port, host)
}
