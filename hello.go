package main

import "fmt"

// must specify the return type of the function, if returning multiple values, use paranthesis
func add(x int) (int, int) {
	return x + x, x * x
}

func main() {
	fmt.Println("Hello, World!")
	a := 23
	fmt.Println(a)

	ptra := &a
	fmt.Println(ptra)

	//var ptrb *int
	if ptra != nil {
		fmt.Println("its not null/nil")
	} else { // the else statement should be right after the closing bracket of the if statement
		fmt.Println("its null/nil")
	}

	var b float32 = 2.324
	fmt.Println(b)

	// a := 123 error as it has already been declared, so just use assignment operator = to reassign the variable

	// after declaring the variable, you must use it or else it will throw an error
	c, d := add(3)
	fmt.Println(c, d)

	var e float32
	fmt.Println(e) // if no value is assigned, it will be assigned the zero value of the type

	e = 3.14

	fmt.Printf("the value of PI is %f", e)

}
