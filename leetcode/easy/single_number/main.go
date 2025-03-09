package single_number

import "sort"

func singleNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	for i := 1; i < len(nums); i += 2 {
		if nums[i] != nums[i-1] {
			return nums[i-1]
		}
	}
	return nums[len(nums)-1]
}

func singleNumberXOR(nums []int) int {
	ans := 0
	for _, num := range nums {
		ans = ans ^ num
	}
	return ans
}
