package permutations

func permute(nums []int) [][]int {
	var dfs func(digits []int)
	res := make([][]int, 0) // why ??
	path := make([]int, 0)

	dfs = func(digits []int) {
		if len(digits) == 0 {
			tmp2 := make([]int, len(path))
			copy(tmp2, path)
			res = append(res, tmp2)
			return
		}
		for i := range digits {
			path = append(path, digits[i])
			tmp := make([]int, len(digits))
			copy(tmp, digits)
			dfs(append(tmp[:i], tmp[i+1:]...)) // remove i-th digit
			path = path[:len(path)-1]          // backtracking
		}
	}
	dfs(nums)
	return res
}
