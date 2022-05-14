package hello

import (
	"math"
)

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectagle struct {
	width  float64
	height float64
}
type Circle struct {
	radius float64
}
type Triagle struct {
	bottom float64
	height float64
}

func (rect Rectagle) Perimeter() float64 {
	return 2 * (rect.height + rect.width)
}

func (rect Rectagle) Area() float64 {
	return rect.height * rect.width
}

func (cir Circle) Perimeter() float64 {
	return 2 * math.Pi * cir.radius
}

func (cir Circle) Area() float64 {
	return math.Pi * cir.radius * cir.radius
}

func (tri Triagle) Perimeter() float64 {
	return 3 * tri.bottom
}
func (tri Triagle) Area() float64 {
	return (tri.bottom * tri.height) / 2
}
