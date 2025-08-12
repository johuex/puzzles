package countingbits

import "math/bits"

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = bits.OnesCount(uint(i))
	}
	return res
}

func countBitsLow(n int) []int {
	res := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// count of 1 of number with shift + odd/even
		res[i] = res[i>>1] + i&1
	}
	return res
}
