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
	retangle := Retangle{10.0, 10.0}
	got := Area(retangle)
	want := 100.0

	if got != want {
		t.Errorf("got %.2f , want %.2f ", got, want)
	}
}
