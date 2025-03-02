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

func missingNumberMath(nums []int) int {
	n := len(nums)
	expectedSum := n * (n + 1) / 2
	var actualSum int
	for _, num := range nums {
		actualSum += num
	}

	return expectedSum - actualSum
}
