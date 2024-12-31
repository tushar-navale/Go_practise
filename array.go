package main

import (
	"fmt"
)

func main() {
	// Array declaration and initialization
	// for array we specify the length of the array
	var arr1 [5]int               // array of 5 integers, initialized to zero
	arr2 := [5]int{1, 2, 3, 4, 5} // array of 5 integers, initialized with values

	// Accessing array elements
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)
	fmt.Println("arr2[2]:", arr2[2]) // accessing the third element

	// Modifying array elements
	arr2[2] = 10
	fmt.Println("Modified arr2:", arr2)

	// Array length
	fmt.Println("Length of arr2:", len(arr2))

	// Slices
	//for slices we dont specify the length of the array
	var slice1 []int               // slice of integers, initially nil
	slice2 := []int{1, 2, 3, 4, 5} // slice of integers, initialized with values

	// Accessing slice elements
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice2[2]:", slice2[2]) // accessing the third element

	// Modifying slice elements
	slice2[2] = 10
	fmt.Println("Modified slice2:", slice2)

	// Slice length and capacity
	fmt.Println("Length of slice2:", len(slice2))
	fmt.Println("Capacity of slice2:", cap(slice2))

	// Creating a slice from an array
	slice3 := arr2[1:4] // slice from index 1 to 3 (excluding index 4)
	fmt.Println("slice3:", slice3)

	// Appending to a slice
	slice3 = append(slice3, 6, 7)
	fmt.Println("Appended slice3:", slice3)

	slice3 = append(slice3, slice2...) // for using another slice we have to specify the ... after the slice name
	fmt.Println("Appended slice3:", slice3)

	// Copying a slice
	// make command will allacte the type of elements specified in the first argument and the second argument will be the length of the slice
	slice4 := make([]int, len(slice3)) // basically an variable array of the same length as slice3

	// copy just copies from 1st para to the 2nd para,
	// if the length of the 2nd para is less than the 1st para,
	// it will only copy the number of elements equal to the length of the 2nd para
	copy(slice4, slice3)
	// if the length of 1st arg is smaller than the 2nd arg error will be thrown

	fmt.Println("Copied slice4:", slice4)

}
