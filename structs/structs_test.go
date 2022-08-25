package structs

import "testing"

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 36.0},
		{shape: Circle{Radius: 10}, want: 62.83185307179586},
		{shape: RightTriangle{Base: 21, Height: 12}, want: 48.874507866387546},
	}

	for _, pt := range perimeterTests {
		got := pt.shape.Perimeter()
		if got != pt.want {
			t.Errorf("got %g, want %g", got, pt.want)
		}
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: RightTriangle{Base: 21, Height: 12}, want: 126.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g, want %g", got, tt.want)
		}
	}
}
