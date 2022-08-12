package structs

import "testing"

func TestPerimeter(t *testing.T) {
	retangle := Retangle{10.0, 10.0}
	got := Perimeter(retangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f , want %.2f ", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Retangle", shape: Retangle{Width: 10, Height: 10}, hasArea: 100},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 10, Height: 10}, hasArea: 50},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.hasArea {
				t.Errorf(" %#v got %g want %g ", tt.shape, got, tt.hasArea)
			}
		})
	}
}
