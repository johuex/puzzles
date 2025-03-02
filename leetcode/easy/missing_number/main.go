package main

func missingNumber(nums []int) int {
	exist := make([]bool, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		exist[i] = false
	}
	for _, num := range nums {
		exist[num] = true
	}
	for idx, val := range exist {
		if !val {
			return idx
		}
	}
	return 0
}
