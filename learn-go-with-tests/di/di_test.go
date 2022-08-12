package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	name := "ed"
	buffer := bytes.Buffer{}
	Greet(&buffer, name)

	got := buffer.String()
	want := "hi~" + name

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
