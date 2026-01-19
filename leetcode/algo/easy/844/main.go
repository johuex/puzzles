package backspacecompare

func cutStr(ss string) string {
	s := []rune(ss)
	idxArr := make([]int, 0)
	// save indexes
	for idx, e := range s {
		if e != '#' {
			idxArr = append(idxArr, idx)
		} else if e == '#' {
			if len(idxArr) != 0 {
				idxArr = idxArr[:len(idxArr)-1] // cut last (previous) elem
			}
		}
	}

	if len(idxArr) == 0 {
		return ""
	}
	// build strings from index arrays
	var idxPtr int
	sc := make([]rune, len(idxArr))
	for idx, e := range s {
		if idx == idxArr[idxPtr] {
			sc[idxPtr] = e
			idxPtr++
			if idxPtr == len(idxArr) {
				break
			}
		}

	}
	return string(sc)
}

func backspaceCompare(s string, t string) bool {
	// save indexes from first raw
	// save indexis from second raw
	// compare string

	sc := cutStr(s)
	tc := cutStr(t)

	if len(sc) != len(tc) {
		return false
	}
	return sc == tc
}
