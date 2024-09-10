package main

import "fmt"

// Define an interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a Rectangle type
type Rectangle struct {
	width, height float64
}

// Implement Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Implement Perimeter method for Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Define a Circle type
type Circle struct {
	radius float64
}

// Implement Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

// Implement Perimeter method for Circle
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}

// A function that takes a Shape interface
func PrintShapeDetails(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	circ := Circle{radius: 7}

	// Both Rectangle and Circle implement Shape, so they can be passed to PrintShapeDetails
	PrintShapeDetails(rect)
	PrintShapeDetails(circ)
}
