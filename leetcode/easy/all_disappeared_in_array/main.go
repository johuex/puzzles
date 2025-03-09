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

func AbsInt(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func findDisappearedNumbersFastest(nums []int) []int {
	for _, num := range nums {
		nums[AbsInt(num)-1] = -AbsInt(nums[AbsInt(num)-1])
	}
	res := make([]int, 0)
	for idx, num := range nums {
		if num > 0 {
			res = append(res, idx+1)
		}
	}
	return res
}
