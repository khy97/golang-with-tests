package structs

import "math"

// Shape -> Interface for shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle -> Struct for Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Perimeter -> Calculates Perimeter of Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area -> Calculates area of Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle -> Struct for Circle
type Circle struct {
	Radius float64
}

// Perimeter -> Calculates Perimeter of Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Area -> Calculates area of Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Triangle -> Struct for Triangle
type Triangle struct {
	Height float64
	Base   float64
}

// Perimeter -> Calculates Perimeter of Triangle
func (t Triangle) Perimeter() float64 {
	return 0
}

// Area -> Calculates area of Triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
