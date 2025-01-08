package main main
import(
	"fmt"
)

func main(){
	user,err := getUser()
	if err != nil {
		fmt.Println(err)
		return
	}
	// use user here
	// in js we would declare all the user and stuff in the try block and error in the catch block
	// and if there is error we would return from the catch block but wouldn't which user function
	// is throwing error, so error handling in JS is much more difficult

	// In go
	// we can easily understand where the error is using the above code we will know where the error is
	// which makes error handling in go much easier

	profile, err := getProfile(user.id)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func getUser() (User, error) {
	// this type of error handling is called explicit error handling, which explicitly 
	// gives which method is throwing error and handling the error is much more robust 
}
