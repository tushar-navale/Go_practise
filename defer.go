package main

import "fmt"

func add1(x int) (int, int) {
	return x + x, x * x
}

func divide(a, b int) float32 {
	if b == 0 {
		panic("Divide by zero error")
	}
	res := float32(a) / float32(b) // convert both a and b to float32 before division
	return res
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r) // it will print the statement that is written in the panic statement which encounters the panic statment
		}
	}()

	defer fmt.Println("hi")  // prints at third
	fmt.Println("first")     // this will be printed first
	defer fmt.Println("Bye") // prints at second

	res := divide(3, 0)
	fmt.Println(res)

	//panic("error occured") // when it encounters this statement then all the deferred statements are executed and then the program stops
	// all the defer and normal statements are excurted before the panic but the deferred are done in a reverse order

	c, d := add1(3)
	fmt.Println(c, d)
}
