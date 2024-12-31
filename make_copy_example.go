package main

import (
	"fmt"
)

func main() {
	// Using make to create a slice
	slice1 := make([]int, 5) // creates a slice of length 5 and capacity 5, initialized to zero
	fmt.Println("slice1:", slice1)
	fmt.Println("Length of slice1:", len(slice1))
	fmt.Println("Capacity of slice1:", cap(slice1))

	// Using make to create a slice with specified length and capacity
	slice2 := make([]int, 3, 5) // creates a slice of length 3 and capacity 5, initialized to zero
	fmt.Println("slice2:", slice2)
	fmt.Println("Length of slice2:", len(slice2))
	fmt.Println("Capacity of slice2:", cap(slice2))

	// Modifying slice elements
	slice2[0] = 1
	slice2[1] = 2
	slice2[2] = 3
	fmt.Println("Modified slice2:", slice2)

	// Using copy to copy elements from one slice to another
	slice3 := []int{4, 5, 6}
	copy(slice1, slice3) // copies elements from slice3 to slice1
	fmt.Println("slice1 after copy:", slice1)
	fmt.Println("slice3:", slice3)

	// Using copy with different lengths
	slice4 := make([]int, 2)
	copy(slice4, slice3) // only the first 2 elements of slice3 are copied to slice4
	fmt.Println("slice4 after copy:", slice4)
}
