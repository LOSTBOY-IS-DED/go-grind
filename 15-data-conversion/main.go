package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 42
	fmt.Println("number is : ", num)
	fmt.Printf("type of num is : %T \n", num)

	var data float64 = float64(num)
	fmt.Println("data is : ", data)
	fmt.Printf("type of data is : %T \n", data)

	// number to string
	num = 123
	var str = strconv.Itoa(num)
	fmt.Println("str is : ", str)
	fmt.Printf("type of str is : %T \n", str)

	// string to number
	str = "123"
	var convertedNum, _ = strconv.Atoi(str)
	fmt.Println("convertedNum is : ", convertedNum)
	fmt.Printf("type of convertedNum is : %T \n", convertedNum)
}
