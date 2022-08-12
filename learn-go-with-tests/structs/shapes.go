package structs

import "math"

type Shape interface {
	Area() float64
}

type Retangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(retangle Retangle) float64 {
	return 2 * (retangle.Width + retangle.Height)
}

func (r Retangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
