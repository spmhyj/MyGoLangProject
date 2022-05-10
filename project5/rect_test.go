package main

import (
	"math"
	"testing"
)

func Test_Perimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("want %.2f but got %.2f", want, got)
		}
	}

	t.Run("Rect Perimeter test", func(t *testing.T) {
		rect := Rectagle{Width: 10.0, Height: 15.0}
		checkPerimeter(t, rect, 50.0)
	})

	t.Run("area Perimeter test", func(t *testing.T) {
		cir := Cirlcle{10.0}
		checkPerimeter(t, cir, math.Pi*10.0)
	})

}

func Test_area(t *testing.T) {
	t.Helper()

	areaTest2 := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "rect", shape:Rectagle{Width: 12.0, Height: 6.0}, want:72.0},
		{"circle", Cirlcle{10}, 314.159},
	}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectagle{12.0, 6.0}, 72.0},
		{Cirlcle{10.0}, 314.1592653589793},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %#v want %#v", got, tt.want)
		}
	}
}
