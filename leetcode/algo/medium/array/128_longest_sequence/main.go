package longestsequence

func longestConsecutive(nums []int) int {
	// Let's say there can be only one sequence in input

	// init map with number: struct{}{} (duplicates aren't counted)
	numsSet := make(map[int]struct{})
	for _, num := range nums {
		numsSet[num] = struct{}{}
	}
	var best int

	// if number is start of sequence (no num-1 in array) -> check sequence by num++
	for num := range numsSet {
		if _, ok := numsSet[num-1]; !ok {
			last := num + 1
			for {
				if _, ok := numsSet[last]; !ok {
					break
				}
				last++
			}
			best = max(best, last-num)
		}
	}

	return best
}
