package majorityelement

func majorityElement(nums []int) int {
	cntMap := make(map[int]int)

	for _, num := range nums {
		cntMap[num] += 1
	}

	lenNums := len(nums)

	for key, val := range cntMap {
		if val > lenNums/2 {
			return key
		}
	}
	return nums[0]
}
