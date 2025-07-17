package main

//NOTE : only for loop is used in this file, no other loops are used
// NOTE : For is the only construct we create while loop using for loop

func main() {
	// while loop fashion
	i := 0
	for i <= 5 {
		println(i)
		i++
	}

	// for loop fashion
	for i := 0; i <= 5; i++ {

		if i == 2 {
			continue
		}

		println(i)
	}

	// infinite loop
	// for {
	// 	println("infinite loop")
	// }

	for i := range 3 {
		println(i) // prints 0, 1, 2
	}
}
