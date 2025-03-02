package main

func findDisappearedNumbers(nums []int) []int {
	unProcessed := make(map[int]struct{})
	for i := 1; i < len(nums)+1; i++ {
		unProcessed[i] = struct{}{}
	}

	for _, num := range nums {
		delete(unProcessed, num)
	}
	res := make([]int, len(unProcessed))
	var idx int
	for key, _ := range unProcessed {
		res[idx] = key
		idx += 1
	}
	return res
}
