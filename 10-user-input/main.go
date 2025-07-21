package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Enter your name:")

	// a wrong wat to read user input coz it breaks after first space
	// var name string
	// fmt.Scanln(&name)
	// fmt.Println("Hello", name, "! Welcome to Go programming.")

	// instead use bufio package
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello", name)
}

//bufio is just a reader which reads from the underlying reader (os.Stdin in this case)
// buffer reader are commonly used to improve the efficiency of reading from a reader
