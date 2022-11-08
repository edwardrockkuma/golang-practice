package main

import (
	m "edward/gopractice/learn-go-with-tests/map_try"
	"fmt"
)

func main() {
	d1 := m.Dictionary{"1": "1"}
	d2 := m.Dictionary{"2": "2"}

	s1, _ := d1.Search("1")
	s2, e2 := d2.Search("1")
	fmt.Printf("s1: %s", s1)
	fmt.Printf("s2: %s,e2: %s", s2, e2.Error())
}
