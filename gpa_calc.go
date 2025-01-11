package main

import (
	"fmt"
)

func main() {
	var total_credits float32 = 0.0
	var total_gpa float32 = 0.0
	var credits float32
	var sgpa float32
	var sem int
	fmt.Print("Enter the number of semesters: ")
	fmt.Scan(&sem)
	for i := 0; i < sem; i++ {
		fmt.Print("Enter the number of number of credits for the sem:", i+1, " ")
		fmt.Scan(&credits)

		total_credits += credits

		fmt.Print("Enter the number of number of sgpa for the sem:", i+1, " ")
		fmt.Scan(&sgpa)

		total_gpa += credits * sgpa
	}

	final_gpa := total_gpa / total_credits

	fmt.Printf("The final gpa is: %.2f", final_gpa)

}
