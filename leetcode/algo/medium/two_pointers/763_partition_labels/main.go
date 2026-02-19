package partitionlabels

func partitionLabels(s string) []int {
	last := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		// collect latest position of each word
		last[s[i]] = i
	}

	res := []int{}
	start, end := 0, 0

	// run over string
	for i := 0; i < len(s); i++ {

		// and found the latest position of first elem or elem inside batch
		if last[s[i]] > end {
			end = last[s[i]]
		}

		// found end, count length of batch
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}

	return res
}
