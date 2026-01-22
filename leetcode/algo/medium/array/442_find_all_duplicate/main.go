package findduplicate

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findDuplicates(nums []int) []int {
	ans := make([]int, 0)

	// we use here init array as "hash" table
	// we mark number under idx with -1 if met it earlier
	// if we meet -1 number again -- duplicate
	for _, num := range nums {
		tmpIdx := Abs(num) - 1
		if nums[tmpIdx] < 0 {
			ans = append(ans, Abs(num))
		} else {
			nums[tmpIdx] *= -1
		}
	}
	return ans
}
