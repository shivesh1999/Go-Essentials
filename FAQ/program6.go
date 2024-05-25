package main

import "fmt"

// Shape is an interface that defines a common behavior for all shapes
type Shape interface {
	Area() float64
}

// Circle is a type that implements the Shape interface
type Circle struct {
	Radius float64
}

// Area calculates the area of a circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Rectangle is a type that implements the Shape interface
type Rectangle struct {
	Width  float64
	Height float64
}

// Area calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 10, Height: 5}

	shapes := []Shape{circle, rectangle}

	var s shape
	s = circle
	fmt.Println("Radius", s.Radius)

	for _, shape := range shapes {
		fmt.Printf("Area of shape: %f\n", shape.Area())
	}
}
