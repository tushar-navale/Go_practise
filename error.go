package main

import (
	"fmt"
	"strconv"
)

// func main(){
// 	user,err := getUser()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	// use user here
// 	// in js we would declare all the user and stuff in the try block and error in the catch block
// 	// and if there is error we would return from the catch block but wouldn't which user function
// 	// is throwing error, so error handling in JS is much more difficult

// 	// In go
// 	// we can easily understand where the error is using the above code we will know where the error is
// 	// which makes error handling in go much easier

// 	profile, err := getProfile(user.id)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// }

// func getUser() (User, error) {
// 	// this type of error handling is called explicit error handling, which explicitly
// 	// gives which method is throwing error and handling the error is much more robust
// }

// error interface is a single method interface with Error() string method
// which returns a string describing the error

func main() {

	a := "343a"
	b, err := str_to_int(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(b)
}

func str_to_int(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("strconv.Atoi: %v, wrong format and operation check the variable", err)
	}
	fmt.Println("%v", err)
	return i, nil
}
