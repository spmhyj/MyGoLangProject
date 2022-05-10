package hello

import (
	"fmt"
	"math"
	"testing"
)

const (
	abc = iota + 10
	def
	ghi
)

func Test_rscHello(t *testing.T) {
	sayHello()
}

func Test_sprintf(t *testing.T) {
	a := 911
	s := fmt.Sprintf("hex [%#x] \t binary [%b] \t %d \t oct [%#o]", a, a, a, a)

	fmt.Println(s)
}

func Test_iota(t *testing.T) {
	fmt.Println(abc, def, ghi)
}

func Test_shapeTable(t *testing.T) {
	t.Helper()
	type ShapeTest struct {
		tShape        Shape
		wantPerimeter float64
		wantArea      float64
	}
	shapeTestList := []ShapeTest{
		{Rectagle{width: 20.0, height: 10}, 60.0, 200.0},
		{Circle{10.0}, 2 * math.Pi * 20.0, math.Pi * 10.0 * 10.0},
		{Triagle{bottom: 10.0, height: 10.0}, 3 * 10.0, (10.0 * 10.0) / 2},
	}

	for _, tt := range shapeTestList {
		fmt.Printf("type [%T]\n", tt.tShape)
		if tt.tShape.Perimeter() != tt.wantPerimeter {
			t.Errorf("Perimeter not matching get [%f] - want [%f]", tt.tShape.Perimeter(), tt.wantPerimeter)
		}
		if tt.tShape.Area() != tt.wantArea {
			t.Errorf("Area not matching get [%f] - want [%f]", tt.tShape.Area(), tt.wantArea)
		}
	}
}
