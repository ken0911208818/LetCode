package PalindromeNumber

import "math"

func reverse1(x int) bool {
	var input int
	var y = 0
	input = x
	if x < 0 {
		input = -x
	}
	for input > 0 {
		r := input % 10
		y = y*10 + r
		input = input / 10
	}
	if y > math.MaxInt32 || y < math.MinInt32 {
		return false
	}
	if x < 0 {
		y = -y
	}
	if y == x {
		return true
	}
	return false
}
