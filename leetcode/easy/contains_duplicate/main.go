package main

func containsDuplicate(nums []int) bool {
	countElems := make(map[int]struct{})
	for _, num := range nums {
		_, ok := countElems[num]
		if ok {
			return true
		} else {
			countElems[num] = struct{}{}
		}
	}
	return false
}
