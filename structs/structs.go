package structs

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type RightTriangle struct {
	Height float64
	Base   float64
}

func (t RightTriangle) Area() float64 {
	return (t.Height * t.Base) / 2
}

func (t RightTriangle) Perimeter() float64 {
	ht := math.Sqrt(t.Height * t.Base)
	return ht + t.Height + t.Base
}
