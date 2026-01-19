package twosum

func twoSum(nums []int, target int) []int {
	idxMap := make(map[int]int)
	for idx, val := range nums {
		idxMap[val] = idx
	}
	for idx, val := range nums {
		if secondIdx, ok := idxMap[target-val]; ok && secondIdx != idx {
			return []int{idx, secondIdx}
		}
	}
	return []int{0, 0}
}
