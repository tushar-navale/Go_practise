package main

import (
	"fmt"
)

func main() {
	// Initial slice
	slice := []int{1, 2, 3}
	fmt.Println("Original slice:", slice)

	// Append a single element
	slice = append(slice, 4)
	fmt.Println("After appending 4:", slice)

	// Append multiple elements
	slice = append(slice, 5, 6, 7)
	fmt.Println("After appending 5, 6, 7:", slice)

	// Append another slice
	anotherSlice := []int{8, 9, 10}
	slice = append(slice, anotherSlice...)
	fmt.Println("After appending another slice:", slice)
}
