package main

import (
	"fmt"
)

// Define an interface
type Shape interface {
	area() float64
	perimeter() float64
}

// Define a struct type
type Circle struct {
	radius float64
}

// Implement the Shape interface for Circle
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// Define another struct type
type Rectangle struct {
	width, height float64
}

// Implement the Shape interface for Rectangle
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func main() {
	// Create instances of Circle and Rectangle
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	// Create a slice of Shape interface
	// as both circle and rectangle implement have the area and perimeter methods it is a part of the shape interface
	shapes := []Shape{circle, rectangle}

	// Iterate over the slice and call the methods
	// here _ is used as we dont need the index, the 1st value is the index and the 2nd value is the shape
	for _, shape := range shapes {
		fmt.Println("Area:", shape.area())
		fmt.Println("Perimeter:", shape.perimeter())
	}
}