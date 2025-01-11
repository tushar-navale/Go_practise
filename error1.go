package main

import (
	"errors"
	"fmt"
)

type User struct {
	name    string
	age     int
	user_id int
}

var (
	count int
	users []User
)

func (u *User) CreateUser() error {
	fmt.Print("Enter the user name: ")
	_, err := fmt.Scan(&u.name)
	if err != nil {
		return errors.New("failed to read user name")
	}

	fmt.Print("Enter the user age: ")
	_, err = fmt.Scan(&u.age)
	if err != nil {
		return errors.New("failed to read user age")
	}

	u.user_id = count + 1
	count++
	users = append(users, *u)
	return nil
}

func GetUser(id int) (User, error) {
	for _, user := range users {
		if user.user_id == id {
			return user, nil
		}
	}
	return User{}, errors.New("user not found")
}

func divide1(n1 int, n2 int) (float32, error) {
	if n2 != 0 {
		res := float32(n1) / float32(n2)
		return res, nil
	}
	return 0, fmt.Errorf("divide by zero error")
}

func main() {
	var u User
	var n int
	fmt.Printf("enter the number of users\n")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("enter the user info of %v \n", i+1)
		err := u.CreateUser()
		if err != nil {
			fmt.Println("Error occurred in creating the user:", err)
			return
		}

		fmt.Println("User created successfully:", u)
	}
	var id int
	fmt.Print("Enter the user ID to retrieve: ")
	fmt.Scan(&id)

	user, err := GetUser(id)
	if err != nil {
		fmt.Println("Error occurred in getting the user:", err)
		return
	}

	fmt.Println("User retrieved successfully:", user)

	fmt.Print("enter 2 value to divide")
	var num1 int
	var num2 int
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	val, err := divide1(num1, num2)

	if err != nil {
		fmt.ErrorF("error %v", err)

	}

	fmt.Println(val)
}
