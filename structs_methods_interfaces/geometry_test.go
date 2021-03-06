package geometry

import "testing"

func TestPerimeter(t *testing.T) {
	esCorrecto := func(t *testing.T, got, expected float64) {
		t.Helper()
		if got != expected {
			t.Errorf("expected %.2f got %.2f", expected, got)
		}
	}
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0
	esCorrecto(t, got, expected)
}
func TestArea(t *testing.T) {
    //TDD Table Driven Development
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}
	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.want)
			}
		})
	}
}
