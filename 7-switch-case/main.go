package main

import "time"

func main() {

	i := 4
	switch i {
	case 1:
		println("one")
	case 2:
		println("two")
	case 3:
		println("three")
	default:
		println("default case")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("It's the weekend!")
	default:
		println("It's a weekday")
	}

	// type switch only in GO

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			println("int")
		case string:
			println("string")
		case bool:
			println("bool")
		default:
			println("unknown type")
		}
	}

	whoAmI(1)
	whoAmI("hello")
}
