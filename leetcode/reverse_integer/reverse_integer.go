package reverseinteger

import (
	"fmt"
	"math"
	"strconv"
)

// convert int to string to solve this challenge
func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	s := strconv.Itoa(x)
	bs := []byte(s)
	head := 0
	tail := len(bs) - 1
	if tail > head {
		bs[head], bs[tail] = bs[tail], bs[head]
		head++
		tail--
	}

	rs := string(bs)

	if sign == -1 {
		rs = "-" + rs
	}

	y, err := strconv.Atoi(rs)
	if err != nil {
		fmt.Println(err, rs)
	}

	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}

	return y
}
