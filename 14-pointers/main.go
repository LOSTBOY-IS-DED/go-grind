package main

import "fmt"

func main() {
	var a = 10
	var pointer = &a

	// num := 20;
	// ptr := &num

	fmt.Println("Location where a is stored : ", pointer) // 0xc0000b9c40 location if a in memory
	fmt.Println("num has value : ", a)
	fmt.Println("data stored in pointer : ", *pointer)

	var anotherPointer *int // default pionter == nil
	if anotherPointer == nil {
		fmt.Println("data is not assigned yet")
	}

	modifyValueByReference(&a) // the modified the data inside the pointer which was passed inside the function
	fmt.Println("num has value : ", a)

}

func modifyValueByReference(value *int) {
	*value = *value * 2
}
