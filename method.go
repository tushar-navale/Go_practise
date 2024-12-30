package main

import (
	"fmt"
)

// Define a struct type
type Rectangle struct {
	width, height int
}

// for multiple line declatration no need for the comma at the end of a object entity
type Employee struct {
	name   string
	age    int
	salary int
}

//embedded struct
type Family struct {
	Employee
	Fam_size int
}


// Method with a receiver of type Rectangle
// here for the area there is no parameter as the object itself is passed as the parameter
func (r Rectangle) area() int {
	return r.width * r.height
}

func (e Employee) display() {
	fmt.Println(e.name, e.age, e.salary)
}

func main() {
	// Create an instance of Rectangle
	rect := Rectangle{width: 10, height: 5}

	// Call the method
	res := rect.area() // here rect is a self reference to the object so in function we dont specify the arguments
	fmt.Println("Area of rectangle:", res)

	// Create an instance of Employee
	emp := Employee{name: "John", age: 25, salary: 50000} // decalation using the key value pair with in {}
	fmt.Println(emp)                                      // prints the object in {} format with no comma separating the values
	fmt.Println(emp.age)                                  // prints the object.age without any {} or comma

	emp1 := Employee{}

	emp1.display()
	emp1.name = "Sam"
	emp1.age = 30
	emp1.salary = 60000
	emp1.display()
	emp1.age = 34
	emp1.display()

	//anonymous strut
	//there is no struct name, just the object is created 
	car1 := struct {
		company string
		year    int
	}{
		company: "Toyota",
		year:    2020,
	}

	fmt.Println(car1)


	// embedded struct
	fam1 := Family{}

	fam1.name = "John"
	fam1.Fam_size = 4
	fam1.age = 52
	fam1.salary = 100000

	fmt.Println(fam1) // prints in {{ will have the employee object},{fam_size}}

}
