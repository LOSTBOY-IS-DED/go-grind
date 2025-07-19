package main

import "fmt"

func main() {
	var nums [4]int

	fmt.Println(len(nums)) // to get the length of the array
	nums[0] = 1            // pushing values into the array
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	var bools [3]bool
	bools[2] = true

	// print complete array
	fmt.Println(bools) // prints [false false true]

	var names [3]string
	names[0] = "Subh"
	fmt.Println(names) // prints [Subh  ]
}
