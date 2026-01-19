package issubsequence

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}
	if t == "" {
		return false
	}

	sp := 0
	tp := 0

	for tp < len(t) {
		if s[sp] == t[tp] {
			sp++
		}
		if sp == len(s) {
			return true
		}
		tp++
	}
	return false
}
