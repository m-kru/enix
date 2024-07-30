package util

import "math"

func NextPowerOfTwo(n int) int {
	return int(math.Pow(2, math.Ceil(math.Log2(float64(n)))))
}
