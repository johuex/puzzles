package longestpalindrome

import (
	"strings"
)

func longestPalindrome(s string) string {
	sMod := s
	sMod = strings.Join(strings.Split(sMod, ""), "#")
	sMod = "#" + sMod + "#"
	// work with real symbols, not bytes
	runes := []rune(sMod)
	n := len(runes)
	palRadArr := make([]int, n)
	var center, rightBoundary int
	// center -- center of th righest palindrome on current index
	// rightBoundary -- center + confirmed radius

	for i := range n {
		// mirroring in range if we in range
		mirror := 2*center - i //  aka left border w/o variable's value overflow with center and radius (i-center)
		if i < rightBoundary {
			// it's main part for reusing calculation
			// There may be the smallest palindromes inside the rightBound
			// just to the left of the center, and we can reuse their values instead of counting them again.
			palRadArr[i] = min(rightBoundary-i, palRadArr[mirror])
		}

		// extending, checking elems left and right from current index
		for i+1+palRadArr[i] < n && i-1-palRadArr[i] >= 0 && runes[i+1+palRadArr[i]] == runes[i-1-palRadArr[i]] {
			palRadArr[i] += 1
		}

		// refresh center and rightBoundary
		// in case if we found equal symbols beyond the pointer boundary (rightBoundary) from current idx
		if i+palRadArr[i] > rightBoundary {
			center = i
			rightBoundary = i + palRadArr[i]
		}
	}

	var maxLength int
	var maxLengthIdx int
	for idx, r := range palRadArr {
		if r > maxLength {
			maxLength = r
			maxLengthIdx = idx
		}
	}
	startIndex := (maxLengthIdx - maxLength) / 2

	return s[startIndex : startIndex+maxLength]
}
