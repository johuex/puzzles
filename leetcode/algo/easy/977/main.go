package squaresorted

import "math"

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sortedSquares(nums []int) []int {
	// set two pointers
	// move right pointer until first positove value or zero
	// select the lowest number by abs under two pointers (move left to left and right to right)

	if len(nums) == 1 {
		return []int{nums[0] * nums[0]}
	}

	var left, right int
	// move right pointer until zero or positive number
	for nums[right] < 0 && right < len(nums)-1 {
		right++
	}
	if right == len(nums)-1 && nums[right-1] < 0 || right == 0 {
		// case when only negative numbers exist
		left = -1
	}
	left = right - 1

	res := make([]int, len(nums))
	var absLeft, absRight int
	for i := range len(nums) {
		if left < 0 {
			absLeft = math.MaxInt
		} else {
			absLeft = Abs(nums[left])
		}

		if right >= len(nums) {
			absRight = math.MaxInt
		} else {
			absRight = Abs(nums[right])
		}

		if absLeft < absRight {
			res[i] = absLeft * absLeft
			left--
		} else {
			res[i] = absRight * absRight
			right++
		}
	}
	return res
}
