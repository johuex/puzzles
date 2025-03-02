package main

func containsDuplicate(nums []int) bool {
	countElems := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := countElems[num]; ok {
			return true
		}
		countElems[num] = struct{}{}
	}
	return false
}
