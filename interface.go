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

type Square struct {
	side float64
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

func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) perimeter() float64 {
	return 4 * s.side
}

func main() {
	// Create instances of Circle and Rectangle
	circle := Circle{radius: 345.33}
	rectangle := Rectangle{width: 10.37, height: 5.343}
	square := Square{side: 12.23}

	// Create a slice of Shape interface
	// as both circle and rectangle implement have the area and perimeter methods it is a part of the shape interface
	shapes := []Shape{circle, rectangle, square}
	// this shape will automatically call the area and perimeter methods of the circle and rectangle
	// as the circle, rectangle and shape are part of the shape interface
	// Iterate over the slice and call the methods
	// here _ is used as we dont need the index, the 1st value is the index and the 2nd value is the shape
	for index, shape := range shapes {
		if index == 0 {
			fmt.Println("Circle")
		} else if index == 1 {
			fmt.Println("Rectangle")
		} else {
			fmt.Println("Square")
		}
		fmt.Println("Area:", shape.area())
		fmt.Println("Perimeter:", shape.perimeter())

	}

	for _, shape := range shapes {
		circle, ok := shape.(Circle)
		if ok {
			fmt.Println("This is a Circle:", circle)
		}

		square, ok := shape.(Square)
		if ok{
			fmt.Println(square)
		}
	}
}
