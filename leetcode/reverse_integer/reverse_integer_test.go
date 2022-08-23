package reverseinteger

import (
	"math"
	"strconv"
	"testing"
)

func TestXxx(t *testing.T) {
	testCases := []struct {
		want int
		got  int
	}{
		{want: 123, got: 321},
		{want: -123, got: -321},
		{want: 120, got: 21},
		{want: math.MaxInt32 + 1, got: 0},
	}

	for i, tc := range testCases {
		t.Run("case"+strconv.Itoa(i), func(t *testing.T) {
			result := reverse(int(tc.want))

			if result != tc.got {
				t.Errorf("result: %v expected: %v", result, tc.got)
			}
		})
	}
}
