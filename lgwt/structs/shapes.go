package structs

import "math"

// interface that defines a "shape"
// super weird you don't have to explicitly define Rectange/Circle to implement Shape
type Shape interface {
	Area() float64
}

// rectange struct and it's methods
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimiter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// circle struct and it's methods
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// triangle struct and it's methods
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
