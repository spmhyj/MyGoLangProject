package main

import (
	"fmt"
	"math"
	"strings"
)

type Rectagle struct {
	Width  float64
	Height float64
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

func (rect Rectagle) Perimeter() float64 {
	return 2 * (rect.Height + rect.Width)
}

func (rect Rectagle) Area() float64 {
	return rect.Height * rect.Width
}

type Cirlcle struct {
	Radius float64
}

func (cir Cirlcle) Perimeter() float64 {
	return (cir.Radius) * math.Pi
}

func (cir Cirlcle) Area() float64 {
	return math.Pi * cir.Radius * cir.Radius
}

func main() {
	fmt.Println("calc_shape testing...")

	mytest()
	// name := "bill"
	// namePointer := &name
	// fmt.Println("A: ", namePointer)
	// fmt.Println("B: ", &namePointer)

	// printPointer(namePointer)
	// var s Shape = Rectagle{10.0, 15.0}

	// fmt.Printf("Rect Perimeter is %.4f\n", s.Perimeter())
	// s = Cirlcle{10.0}
	// fmt.Printf("circle Perimeter is %.4f\n", s.Perimeter())

}

func mytest() {
	var a []string
	type dock []string

	a = strings.Split("this is a test", " ")
	fmt.Printf("%+v", dock(a))
}
func printPointer(namePointer *string) {
	fmt.Println("C: ", namePointer)
	fmt.Println("D: ", &namePointer)
}
