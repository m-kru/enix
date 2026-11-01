package util

import "math/bits"

// NextPowerOfTwo rounds n up to a power of two, returning 0 for n <= 0.
// The caller must ensure the resulting power of two fits in an int.
func NextPowerOfTwo(n int) int {
	if n <= 0 {
		return 0
	}
	return 1 << bits.Len(uint(n-1))
}
