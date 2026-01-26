package lettercasepermutation

import "unicode"

func letterCasePermutation(s string) []string {
	res := []string{}
	runes := []rune(s) // work with symbols, not bytes
	path := make([]rune, len(runes))

	// defined function in a function to see `res` variable w/o passing to func
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(runes) {
			// one of permutation done, save result
			res = append(res, string(path))
			return
		}

		char := runes[i]

		if unicode.IsLetter(char) {
			// two branch: lowercase or uppercase char
			path[i] = unicode.ToLower(char)
			dfs(i + 1)

			path[i] = unicode.ToUpper(char)
			dfs(i + 1)
		} else {
			// can't uppercase digital
			// so have only one path
			path[i] = char
			dfs(i + 1)
		}
	}

	dfs(0)
	return res
}
