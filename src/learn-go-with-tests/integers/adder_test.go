package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	expected := 3

	if result != expected {
		t.Errorf("expected %d but got %d ", expected, result)
	}
}

// example function will not be executed if remove the 'Output' comment
func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	// Output: 3
}
